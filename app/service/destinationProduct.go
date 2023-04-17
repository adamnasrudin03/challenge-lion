package service

import (
	"adamnasrudin03/challenge-lion/app/entity"
	"adamnasrudin03/challenge-lion/app/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DestinationProductService interface {
	GetAll(ctx *gin.Context) (result []entity.DestinationProduct, statusCode int, err error)
}

type dpSrv struct {
	Repo              repository.DestinationProductRepository
	SourceProductRepo repository.SourceProductRepository
}

func NewDestinationProductService(repo repository.DestinationProductRepository, spRepo repository.SourceProductRepository) DestinationProductService {
	return &dpSrv{
		Repo:              repo,
		SourceProductRepo: spRepo,
	}
}

func (srv *dpSrv) GetAll(ctx *gin.Context) (result []entity.DestinationProduct, statusCode int, err error) {
	result, err = srv.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("[DestinationProductService-GetAll] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	source, err := srv.SourceProductRepo.GetAll(ctx)
	if err != nil {
		log.Printf("[DestinationProductService-GetAll] error get data Source repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	for _, v := range result {
		go func(dp entity.DestinationProduct) {
			sp := entity.SourceProduct{
				ID:           dp.ID,
				ProductName:  dp.ProductName,
				Qty:          dp.Qty,
				PromoPrice:   dp.PromoPrice,
				SellingPrice: dp.SellingPrice,
			}
			_, _ = srv.SourceProductRepo.UpdateByID(ctx, dp.ID, sp)

		}(v)
	}

	for _, v := range source {
		go func(sp entity.SourceProduct) {
			ds := entity.DestinationProduct{
				ID:           sp.ID,
				ProductName:  sp.ProductName,
				Qty:          sp.Qty,
				PromoPrice:   sp.PromoPrice,
				SellingPrice: sp.SellingPrice,
			}
			_, _ = srv.Repo.UpdateByID(ctx, sp.ID, ds)

		}(v)
	}
	return result, http.StatusOK, nil
}
