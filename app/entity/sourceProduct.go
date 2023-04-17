package entity

import "time"

type SourceProduct struct {
	ID           uint64    `gorm:"primaryKey" json:"id"`
	ProductName  string    `gorm:"not null" json:"product_name"`
	Qty          uint64    `json:"qty" `
	SellingPrice uint64    `json:"selling_price" `
	PromoPrice   uint64    `json:"promo_price" `
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (SourceProduct) TableName() string {
	return "source_product"
}
