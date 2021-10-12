package main

import (
	"github.com/gamberooni/go-cats/model"
	"github.com/gamberooni/go-cats/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// create sqlite db instance
	db, err := gorm.Open(sqlite.Open("cats.db"), &gorm.Config{Logger: &logger.Recorder})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Cat{})

	// create router using the sqlite db
	r := router.New(db)

	// start server
	r.Logger.Fatal(r.Start(":80"))
}
