package seeders

import (
	"adamnasrudin03/challenge-lion/app/entity"

	"gorm.io/gorm"
)

func InsertDBOne(db *gorm.DB) {
	tx := db.Begin()
	db.Exec("DELETE FROM source_product")
	var sourceProducts *[]entity.SourceProduct

	db.Find(&sourceProducts)
	if len(*sourceProducts) <= 0 {
		sourceProducts = SeedSourceProducts()
		tx.Create(sourceProducts)
	}

	tx.Commit()
}

func InsertDBTwo(db *gorm.DB) {
	tx := db.Begin()
	db.Exec("DELETE FROM destination_product")
	var destinationProducts *[]entity.DestinationProduct

	db.Find(&destinationProducts)
	if len(*destinationProducts) <= 0 {
		destinationProducts = SeedDestinationProducts()
		tx.Create(destinationProducts)
	}

	tx.Commit()
}
