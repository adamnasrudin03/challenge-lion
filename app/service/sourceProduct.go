package service

import (
	"adamnasrudin03/challenge-lion/app/entity"
	"adamnasrudin03/challenge-lion/app/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SourceProductService interface {
	GetAll(ctx *gin.Context) (result []entity.SourceProduct, statusCode int, err error)
}

type spSrv struct {
	Repo                    repository.SourceProductRepository
	DestionationProductRepo repository.DestinationProductRepository
}

func NewSourceProductService(repo repository.SourceProductRepository, dpRepo repository.DestinationProductRepository) SourceProductService {
	return &spSrv{
		Repo:                    repo,
		DestionationProductRepo: dpRepo,
	}
}

func (srv *spSrv) GetAll(ctx *gin.Context) (result []entity.SourceProduct, statusCode int, err error) {
	result, err = srv.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("[SourceProductService-GetAll] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	destination, err := srv.DestionationProductRepo.GetAll(ctx)
	if err != nil {
		log.Printf("[SourceProductService-GetAll] error get data Destination repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	for _, v := range result {
		go func(sp entity.SourceProduct) {
			dp := entity.DestinationProduct{
				ID:           sp.ID,
				ProductName:  sp.ProductName,
				Qty:          sp.Qty,
				PromoPrice:   sp.PromoPrice,
				SellingPrice: sp.SellingPrice,
			}
			_, _ = srv.DestionationProductRepo.UpdateByID(ctx, sp.ID, dp)

		}(v)
	}

	for _, v := range destination {
		go func(dp entity.DestinationProduct) {
			source := entity.SourceProduct{
				ID:           dp.ID,
				ProductName:  dp.ProductName,
				Qty:          dp.Qty,
				PromoPrice:   dp.PromoPrice,
				SellingPrice: dp.SellingPrice,
			}
			_, _ = srv.Repo.UpdateByID(ctx, dp.ID, source)

		}(v)
	}

	return result, http.StatusOK, nil
}
