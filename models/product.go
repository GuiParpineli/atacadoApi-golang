package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;not null;" json:"name"`
	ProviderName  string `gorm:"size:255;not null;" json:"providerName"`
	Ncm           string `gorm:"size:255;not null;" json:"ncm"`
	Category      string `gorm:"size:255;not null; unique" json:"category"`
	Inventory     int    `gorm:"not null;" json:"inventory"`
	PurchasePrice uint   `gorm:"not null;" json:"purchasePrice"`
	SalePrice     uint   `gorm:"not null;" json:"salePrice"`
}
