package cmd

import (
	  "github.com/urfave/cli/v2"
    // "fmt"
		"log"
		// "time"
    "sync"
    "github.com/valyala/fasthttp"
)

func dbCommand(ctx *cli.Context) error{
	db, err := gorm.Open("sqlite3", "mydb.sqlite3")
	 if err != nil {
			 log.Fatal(err)
	 }
	 if err = db.DB().Ping(); err != nil {
			 log.Fatal(err)
	 }

	 db.LogMode(true)

	 m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
			 {
					 ID: "201608301400",
					 Migrate: func(tx *gorm.DB) error {
							 return tx.AutoMigrate(&Person{}).Error
					 },
					 Rollback: func(tx *gorm.DB) error {
							 return tx.DropTable("people").Error
					 },
			 },
			 {
					 ID: "201608301430",
					 Migrate: func(tx *gorm.DB) error {
							 return tx.AutoMigrate(&Pet{}).Error
					 },
					 Rollback: func(tx *gorm.DB) error {
							 return tx.DropTable("pets").Error
					 },
			 },
	 })

	 err = m.Migrate()
	 if err == nil {
			 log.Printf("Migration did run successfully")
	 } else {
			 log.Printf("Could not migrate: %v", err)
	 }
}
