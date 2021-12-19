package viewer

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/ktakenaka/gosample2022/app/models"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/testsupport/factory"
	"github.com/ktakenaka/gosample2022/testsupport/testdb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func Test_sampleViewer_List(t *testing.T) {
	t.Run("when samples exist", func(t *testing.T) {
		read, write, release := testdb.GetFuncs()
		defer release()

		writeDB, ctx := write(), context.Background()
		user, office, ou := factory.MustBuildOfficeUser(nil)
		_ = user.Upsert(ctx, writeDB, boil.Infer(), boil.Infer())
		_ = office.Upsert(ctx, writeDB, boil.Infer(), boil.Infer())
		_ = ou.Upsert(ctx, writeDB, boil.Infer(), boil.Infer())

		id, _ := ulid.GenerateID()
		sample := &models.Sample{
			ID:        id,
			Title:     "title",
			Category:  "small",
			Date:      time.Now(),
			Amount:    types.NewDecimal(decimal.New(123, 3)),
			CreatedBy: user.ID,
			UpdatedBy: user.ID,
		}

		_ = office.Reload(ctx, writeDB)
		_ = office.AddSamples(ctx, writeDB, true, sample)

		viewer := NewSampleViewer(read)
		samples, _ := viewer.List(ctx)

		if !bytes.Equal(samples[0].ID, sample.ID) {
			t.Errorf("want: %v, got: %v", sample, samples)
		}
	})
}
