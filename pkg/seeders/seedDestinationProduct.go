package seeders

import (
	"adamnasrudin03/challenge-lion/app/entity"
	"fmt"
)

func SeedDestinationProducts() *[]entity.DestinationProduct {
	products := []entity.DestinationProduct{}

	for i := 1; i <= 500; i++ {
		product := entity.DestinationProduct{
			ID:           uint64(i),
			ProductName:  fmt.Sprintf("Product %v", i),
			Qty:          0,
			SellingPrice: 0,
			PromoPrice:   0,
		}
		products = append(products, product)
	}

	return &products
}
