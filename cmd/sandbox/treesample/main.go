package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/app/interface/infrastructure"
	"github.com/ktakenaka/gosample2022/app/pkg/historydate"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	cfg, _         = config.Initialize()
	ctx            = context.Background()
	validFrom, _   = historydate.ParseDate("2021-08-02")
	maxDate, _     = historydate.ParseDate("9999-12-31")
	maxDateTime, _ = time.Parse("2006-01-02 15:04:05", "9999-12-31 23:59:59")
)

func init() {
	boil.DebugMode = true
}

func main() {
	db, task, _ := mysql.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	office := models.Offices().OneP(ctx, db)
	biid := ulid.MustNew()
	sample := &models.TreeSample{
		Biid:            biid,
		Path:            biid,
		OfficeID:        office.ID,
		Name:            "name",
		ValidFrom:       time.Now(),
		ValidTo:         maxDate.ToTime(),
		TransactionFrom: time.Now(),
		TransactionTo:   maxDateTime,
	}
	sample.InsertP(ctx, db, boil.Infer())

	child1 := addTo(db, sample, &models.TreeSample{Biid: ulid.MustNew(), Name: "name2"})
	_ = addTo(db, child1, &models.TreeSample{Biid: ulid.MustNew(), Name: "name3"})

	show(db, biid)
	delete(db, sample)
}

func addTo(db infrastructure.DB, parent, newChild *models.TreeSample) *models.TreeSample {
	fmt.Println("=== addTo ===")
	newChild.Path = constractPath(parent.Path, newChild.Biid)
	newChild.OfficeID = parent.OfficeID

	// TODO: Assign appropriate value
	newChild.ValidTo, newChild.ValidFrom = maxDate.ToTime(), parent.ValidFrom
	newChild.TransactionFrom, newChild.TransactionTo = time.Now(), maxDateTime

	newChild.InsertP(ctx, db, boil.Infer())
	return newChild
}

func show(db infrastructure.DB, biid string) {
	fmt.Println("=== show ===")
	sample := models.TreeSamples(
		// TODO: Add `ValidAt` and `TransactionAt`
		models.TreeSampleWhere.Biid.EQ(biid),
	).OneP(ctx, db)
	children := models.TreeSamples(
		// TODO: Add `ValidAt` and `TransactionAt`
		qm.Where("path LIKE ?", sample.Path+"%"),
	).AllP(ctx, db)

	fmt.Printf("sample: %v\n", sample)
	for i, child := range children {
		fmt.Printf("child%d: %v\n", i, child)
	}
}

func delete(db infrastructure.DB, sample *models.TreeSample) {
	fmt.Println("=== delete ===")
	children := models.TreeSamples(
		// TODO: Add `ValidAt` and `TransactionAt`
		qm.Where("path LIKE ?", sample.Path+"%"),
	).AllP(ctx, db)

	var (
		validTo       = sample.ValidFrom.AddDate(0, 0, 10) // TODO: Use correct one
		transactionTo = time.Now()
	)

	var inserting, updating models.TreeSampleSlice

	// Copy and assign new values
	newSample := *sample
	newSample.ID = 0
	newSample.ValidTo = validTo
	newSample.TransactionTo = maxDateTime
	inserting = append(inserting, &newSample)

	// Revoke out-dated record
	sample.TransactionTo = transactionTo
	updating = append(updating, sample)

	for _, s := range children {
		// Copy and assign new values
		new := *s
		new.ID = 0
		new.ValidTo = validTo
		new.TransactionTo = maxDateTime
		inserting = append(inserting, &new)

		// Revoke out-dated record
		s.TransactionTo = transactionTo
		updating = append(updating, s)
	}

	for _, s := range inserting {
		s.InsertP(ctx, db, boil.Infer())
	}
	for _, s := range updating {
		s.UpdateP(ctx, db, boil.Infer())
	}
}

func constractPath(path ...string) string {
	return strings.Join(path, "/")
}
