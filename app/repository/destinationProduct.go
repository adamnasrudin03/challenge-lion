package repository

import (
	"log"

	"adamnasrudin03/challenge-lion/app/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DestinationProductRepository interface {
	GetAll(ctx *gin.Context) (result []entity.DestinationProduct, err error)
	UpdateByID(ctx *gin.Context, ID uint64, input entity.DestinationProduct) (result entity.DestinationProduct, err error)
}

type dpRepo struct {
	DB *gorm.DB
}

func NewDestinationProductRepository(db *gorm.DB) DestinationProductRepository {
	return &dpRepo{
		DB: db,
	}
}

func (repo *dpRepo) GetAll(ctx *gin.Context) (result []entity.DestinationProduct, err error) {
	query := repo.DB.WithContext(ctx)

	err = query.Find(&result).Error
	if err != nil {
		log.Printf("[DestinationProductRepository-GetAll] error get data: %+v \n", err)
		return
	}

	return
}

func (repo *dpRepo) UpdateByID(ctx *gin.Context, ID uint64, input entity.DestinationProduct) (result entity.DestinationProduct, err error) {
	query := repo.DB.WithContext(ctx)

	log.Printf("Process update Destination Product, id: %+v \n", ID)
	err = query.Model(&result).Where("id = ?", ID).Updates(input).Error
	if err != nil {
		log.Printf("[DestinationProductRepository-UpdateByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	return result, err
}
