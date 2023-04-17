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
func SetupDBonnectionTwo() *gorm.DB {
	configs := configs.GetInstance()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.Dbconfigtwo.Host,
		configs.Dbconfigtwo.Username,
		configs.Dbconfigtwo.Password,
		configs.Dbconfigtwo.Dbname,
		configs.Dbconfigtwo.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	if configs.Dbconfigtwo.DebugMode {
		db = db.Debug()
	}

	if configs.Dbconfigtwo.DbIsMigrate {
		//auto migration entity db
		db.AutoMigrate(
			entity.DestinationProduct{},
		)
	}
	// async insert record to db
	go seeders.InsertDBTwo(db)

	log.Println("Connection Database Success!")
	return db
}

// Close Db Connection method is closing a connection between your app and your db
func CloseDBConnectioTwo(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}
