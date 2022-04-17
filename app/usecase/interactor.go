package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/transaction"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Interactor interface {
	OfficeOne(ctx context.Context, officeID string) (*Office, error)

	SampleList(ctx context.Context, office *Office) (models.SampleSlice, error)
	SampleCreate(ctx context.Context, office *Office, req *BiTemporalSampleRequest) (*Sample, error)
	SampleAddFirst(ctx context.Context, office *Office, req *BiTemporalSampleRequest) error

	SyncSamples(ctx context.Context, samples []*SampleCopy) error
}

type interactor struct {
	p *infrastructure.Provider
}

func NewInteractor(p *infrastructure.Provider) Interactor {
	return &interactor{p: p}
}

func (i *interactor) OfficeOne(ctx context.Context, officeID string) (*Office, error) {
	office, err := models.FindOffice(ctx, i.p.DB, officeID)
	if err != nil {
		return nil, err
	}

	return &Office{Office: office}, nil
}

func (i *interactor) SampleList(ctx context.Context, office *Office) (models.SampleSlice, error) {
	return office.Samples().All(ctx, i.p.DB)
}

/*
TODO: Implement BiTemporalDataModel logics seriously
*/
func (i *interactor) SampleCreate(ctx context.Context, office *Office, req *BiTemporalSampleRequest) (*Sample, error) {
	err := office.AddSamples(
		ctx, i.p.DB, true,
		&models.Sample{
			Biid:      ulid.MustNew(),
			Code:      req.Code,
			Category:  req.Category,
			Amount:    req.Amount,
			ValidFrom: req.ValidFrom.ToTime(),
			ValidTo:   req.ValidTo.ToTime(), // TODO: When validTo is nil, make it max date.
			Version:   1,
		},
	)
	if err != nil {
		return nil, err
	}
	return &Sample{Sample: office.R.Samples[0]}, nil
}
func (i *interactor) SampleAddFirst(ctx context.Context, office *Office, req *BiTemporalSampleRequest) error {
	latest, err := office.Samples(
		models.SampleWhere.Biid.EQ(req.Biid),
		qm.OrderBy(models.SampleTableColumns.ValidTo+" DESC"),
	).One(ctx, i.p.DB)
	if err != nil {
		return err
	}

	/*
		TODO: validation
	*/

	err = transaction.TxExecute(ctx, i.p.DB, func(tx *sql.Tx) error {
		latest.DeletedAt = null.TimeFrom(time.Now())
		latest.Version += 1
		if _, err := latest.Update(ctx, tx, boil.Whitelist(models.SampleCopyColumns.DeletedAt, models.SampleCopyColumns.Version)); err != nil {
			return err
		}

		new := latest
		new.ID = 0
		new.DeletedAt.Valid = false
		new.ValidTo = req.ValidFrom.ToTime().AddDate(0, 0, -1)
		new.Version = 1
		err = office.AddSamples(
			ctx, tx, true,
			new,
			&models.Sample{
				Biid:      new.Biid,
				Code:      new.Code,
				Category:  req.Category,
				Amount:    req.Amount,
				ValidFrom: req.ValidFrom.ToTime(),
				ValidTo:   req.ValidTo.ToTime(),
				Version:   1,
			},
		)
		return err
	})

	return err
}

func (i *interactor) SyncSamples(ctx context.Context, samples []*SampleCopy) (err error) {
	isExist, err := models.Offices(models.OfficeWhere.ID.EQ(samples[0].OfficeID)).Exists(ctx, i.p.DB)
	if err != nil {
		return err
	}
	if !isExist {
		return nil
	}

	sampleIDs := make([]uint, len(samples))
	for i := range samples {
		sampleIDs[i] = samples[i].ID
	}

	err = transaction.TxExecute(ctx, i.p.DB, func(tx *sql.Tx) error {
		existingSamples, err := models.SampleCopies(
			models.SampleCopyWhere.ID.IN(sampleIDs),
			qm.WithDeleted(),
			qm.For("UPDATE"),
		).All(ctx, tx)
		if err != nil {
			return err
		}

		existingSamplesMap := make(map[uint]*models.SampleCopy)
		for i := range existingSamples {
			existingSamplesMap[existingSamples[i].ID] = existingSamples[i]
		}

		var upsertingList models.SampleCopySlice
		for i := range samples {
			existing, ok := existingSamplesMap[samples[i].ID]
			if !ok {
				upsertingList = append(upsertingList, samples[i].SampleCopy)
				continue
			}

			if existing.Version >= samples[i].Version {
				continue
			}

			if diff := cmp.Diff(samples[i].SampleCopy, existing); diff == "" {
				continue
			}

			upsertingList = append(upsertingList, samples[i].SampleCopy)
		}

		// TODO: Use UpsertAll
		for i := range upsertingList {
			if err := upsertingList[i].Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
