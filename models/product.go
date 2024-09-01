package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name         string  `gorm:"size:155;not null;" json:"name"`
	Description  string  `gorm:"size:255;not null;" json:"description"`
	Sku          string  `gorm:"size:155;not null; unique" json:"sku"`
	Barcode      string  `gorm:"size:155;not null; unique" json:"barcode"`
	ProviderName string  `gorm:"size:155;not null;" json:"providerName"`
	Ncm          int     `json:"ncm"`
	Category     string  `gorm:"size:50;not null; unique" json:"category"`
	Inventory    int     `gorm:"not null;" json:"inventory"`
	Price        float64 `gorm:"not null;" json:"price"`
	Cost         float64 `gorm:"not null;" json:"cost"`
}
