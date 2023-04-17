package repository

import (
	"log"

	"adamnasrudin03/challenge-lion/app/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SourceProductRepository interface {
	GetAll(ctx *gin.Context) (result []entity.SourceProduct, err error)
	UpdateByID(ctx *gin.Context, ID uint64, input entity.SourceProduct) (result entity.SourceProduct, err error)
}

type spRepo struct {
	DB *gorm.DB
}

func NewSourceProductRepository(db *gorm.DB) SourceProductRepository {
	return &spRepo{
		DB: db,
	}
}

func (repo *spRepo) GetAll(ctx *gin.Context) (result []entity.SourceProduct, err error) {
	query := repo.DB.WithContext(ctx)

	err = query.Find(&result).Error
	if err != nil {
		log.Printf("[SourceProductRepository-GetAll] error get data: %+v \n", err)
		return
	}

	return
}

func (repo *spRepo) UpdateByID(ctx *gin.Context, ID uint64, input entity.SourceProduct) (result entity.SourceProduct, err error) {
	query := repo.DB.WithContext(ctx)

	err = query.Clauses(clause.Returning{}).Model(&result).Where("id = ?", ID).Updates(input).Error
	if err != nil {
		log.Printf("[SourceProductRepository-UpdateByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	return result, err
}
