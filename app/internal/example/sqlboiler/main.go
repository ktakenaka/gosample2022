package main

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/models"
)

func main() {
	db, err := sqlx.Connect("mysql", "gomsx:gomsx@tcp(database:3306)/gomsx?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	boil.DebugMode = true

	ctx := context.Background()
	pilots, err := models.Pilots().All(ctx, db)
	log.Println(pilots)

	p := models.Pilot{Name: "hello"}
	p.Insert(ctx, db, boil.Infer())

	pilots, err = models.Pilots().All(ctx, db)
	log.Println(pilots[0].ID)
}
