package app

import (
	"adamnasrudin03/challenge-lion/app/repository"
	"adamnasrudin03/challenge-lion/app/service"

	"gorm.io/gorm"
)

func WiringRepository(dbOne *gorm.DB, dbTwo *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		SourceProduct:      repository.NewSourceProductRepository(dbOne),
		DestinationProduct: repository.NewDestinationProductRepository(dbTwo),
	}
}

func WiringService(repo *repository.Repositories) *service.Services {
	return &service.Services{
		SourceProduct:      service.NewSourceProductService(repo.SourceProduct, repo.DestinationProduct),
		DestinationProduct: service.NewDestinationProductService(repo.DestinationProduct, repo.SourceProduct),
	}
}
