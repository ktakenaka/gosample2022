package usecase

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Interactor interface {
	OfficeOne(ctx context.Context, officeID string) (*models.Office, error)

	SampleList(ctx context.Context, office *models.Office) (models.SampleSlice, error)
	SampleCreate(ctx context.Context, office *models.Office, req *BiTemporalSampleRequest) (*models.Sample, error)
	SampleAddFirst(ctx context.Context, office *models.Office, req *BiTemporalSampleRequest) error
}

type interactor struct {
	p *infrastructure.Provider
}

func NewInteractor(p *infrastructure.Provider) Interactor {
	return &interactor{p: p}
}

func (i *interactor) OfficeOne(ctx context.Context, officeID string) (*models.Office, error) {
	return models.Offices(models.OfficeWhere.ID.EQ(officeID)).One(ctx, i.p.DB)
}

func (i *interactor) SampleList(ctx context.Context, office *models.Office) (models.SampleSlice, error) {
	return office.Samples().All(ctx, i.p.DB)
}

/*
TODO: Implement BiTemporalDataModel logics seriously
*/
func (i *interactor) SampleCreate(ctx context.Context, office *models.Office, req *BiTemporalSampleRequest) (*models.Sample, error) {
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
	return office.R.Samples[0], nil
}
func (i *interactor) SampleAddFirst(ctx context.Context, office *models.Office, req *BiTemporalSampleRequest) error {
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
func (i *interactor) SampleAddMiddle(ctx context.Context, office *models.Office) error {
	return nil
}
func (i *interactor) SampleAbolish(ctx context.Context, office *models.Office) error {
	return nil
}
