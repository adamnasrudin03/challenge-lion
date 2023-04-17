package database

import (
	"fmt"
	"log"

	"adamnasrudin03/challenge-lion/app/configs"
	"adamnasrudin03/challenge-lion/app/entity"
	"adamnasrudin03/challenge-lion/pkg/seeders"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Setup Db Connection is creating a new connection to our database
func SetupDBonnectionOne() *gorm.DB {
	configs := configs.GetInstance()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.Dbconfigone.Host,
		configs.Dbconfigone.Username,
		configs.Dbconfigone.Password,
		configs.Dbconfigone.Dbname,
		configs.Dbconfigone.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	if configs.Dbconfigone.DebugMode {
		db = db.Debug()
	}

	if configs.Dbconfigone.DbIsMigrate {
		//auto migration entity db
		db.AutoMigrate(
			entity.SourceProduct{},
		)
	}
	// async insert record to db
	go seeders.InsertDBOne(db)

	log.Println("Connection Database Success!")
	return db
}

// Close Db Connection method is closing a connection between your app and your db
func CloseDBConnectionOne(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}
