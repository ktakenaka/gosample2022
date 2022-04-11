package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/go-cmp/cmp"
	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/debeziumcsmr"
	"github.com/ktakenaka/gosample2022/app/pkg/transaction"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Interactor interface {
	OfficeOne(ctx context.Context, officeID string) (*Office, error)

	SampleList(ctx context.Context, office *Office) (models.SampleSlice, error)
	SampleCreate(ctx context.Context, office *Office, req *BiTemporalSampleRequest) (*Sample, error)
	SampleAddFirst(ctx context.Context, office *Office, req *BiTemporalSampleRequest) error

	SyncSamples(ctx context.Context, tID string, samples []*SampleCopy) error
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

	tx, err := i.p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	if _, err := latest.Delete(ctx, tx, false); err != nil {
		return err
	}

	latest.ID = 0
	latest.DeletedAt.Valid = false
	latest.ValidTo = req.ValidFrom.ToTime().AddDate(0, 0, -1)
	if err := office.AddSamples(
		ctx, tx, true,
		latest,
		&models.Sample{
			Biid:      latest.Biid,
			Code:      latest.Code,
			Category:  req.Category,
			Amount:    req.Amount,
			ValidFrom: req.ValidFrom.ToTime(),
			ValidTo:   req.ValidTo.ToTime(),
		},
	); err != nil {
		return err
	}

	return nil
}

func (i *interactor) SyncSamples(ctx context.Context, tID string, samples []*SampleCopy) (err error) {
	isExist, err := models.Offices(models.OfficeWhere.ID.EQ(samples[0].OfficeID)).Exists(ctx, i.p.DB)
	if err != nil {
		return err
	}
	if !isExist {
		return nil
	}

	err = transaction.TxExecute(ctx, i.p.DB, func(tx *sql.Tx) error {
		var upsertingList models.SampleCopySlice
		for i := range samples {
			// TODO: あーーこれ一斉にロック取らんとデッドロックするやん。
			// ちゃんと楽観的ロック実装しよう
			existing, err := models.SampleCopies(
				models.SampleCopyWhere.ID.EQ(samples[i].ID),
				qm.WithDeleted(),
				qm.For("UPDATE"),
			).One(ctx, tx)
			if err != nil && errors.Is(err, sql.ErrNoRows) {
				upsertingList = append(upsertingList, samples[i].SampleCopy)
				continue
			}
			if err != nil {
				return err
			}

			// TODO: Use version column to judge if we should update
			if existing.DeletedAt.Valid {
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
	if err != nil {
		return err
	}

	return i.p.Redis.Del(ctx, debeziumcsmr.RedisKeyCount(tID), debeziumcsmr.RedisKeyRecords(tID)).Err()
}
