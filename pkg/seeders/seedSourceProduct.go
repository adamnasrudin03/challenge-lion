package seeders

import (
	"adamnasrudin03/challenge-lion/app/entity"
	"fmt"
	"math/rand"
)

func SeedSourceProducts() *[]entity.SourceProduct {
	products := []entity.SourceProduct{}

	for i := 1; i <= 500; i++ {
		product := entity.SourceProduct{
			ID:           uint64(i),
			ProductName:  fmt.Sprintf("Source Product %v", i),
			Qty:          uint64(rand.Intn(100) + 1),
			SellingPrice: 0,
			PromoPrice:   uint64(rand.Intn(100) + 1),
		}
		products = append(products, product)
	}
	return &products
}
