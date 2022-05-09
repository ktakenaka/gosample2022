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
	cfg, _       = config.Initialize()
	ctx          = context.Background()
	validFrom, _ = historydate.ParseDate("2021-08-02")
	maxDate, _   = historydate.ParseDate("9999-12-31")

	blackList = boil.Blacklist(models.TreeSampleColumns.ID, models.TreeSampleColumns.ValidTo, models.TreeSampleColumns.TransactionTo)
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
		TransactionFrom: time.Now(),
	}
	sample.InsertP(ctx, db, blackList)

	child1 := addTo(db, sample, &models.TreeSample{Biid: ulid.MustNew(), Name: "name2"})
	_ = addTo(db, child1, &models.TreeSample{Biid: ulid.MustNew(), Name: "name3"})

	show(db, biid)
	delete(db, child1)
}

func addTo(db infrastructure.DB, parent, newChild *models.TreeSample) *models.TreeSample {
	fmt.Println("=== addTo ===")
	newChild.Path = constractPath(parent.Biid, newChild.Biid)
	newChild.OfficeID = parent.OfficeID

	// FIXME: Assign appropriate value
	newChild.ValidFrom = parent.ValidFrom
	newChild.TransactionFrom = time.Now()
	newChild.InsertP(ctx, db, blackList)
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
		qm.Where("path LIKE ", "%"+biid),
	).AllP(ctx, db)

	fmt.Printf("sample: %v\n", sample)
	fmt.Printf("children: %v\n", children)
}

func delete(db infrastructure.DB, sample *models.TreeSample) {
	fmt.Println("=== delete ===")
	children := models.TreeSamples(
		// TODO: Add `ValidAt` and `TransactionAt`
		qm.Where("path LIKE ", "%"+sample.Biid),
	).AllP(ctx, db)

	validTo := sample.ValidFrom.AddDate(0, 0, 10) // TODO: correct one
	transactionTo := time.Now()
	var upserting models.TreeSampleSlice
	newSample := *sample
	sample.TransactionTo = transactionTo
	newSample.ID = 0
	newSample.ValidTo = validTo
	upserting = append(upserting, sample, &newSample)

	for _, s := range children {
		new := *s
		s.TransactionTo = transactionTo
		new.ID = 0
		new.ValidTo = validTo
		upserting = append(upserting, s, &new)
	}

	// upserting.UpsertAll
}

func constractPath(path ...string) string {
	return strings.Join(path, "/")
}
