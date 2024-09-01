package models

import "github.com/jinzhu/gorm"

type ProductOrder struct {
	gorm.Model
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,constraint:OnDelete:CASCADE;"`
	ProductID uint    `gorm:"not null;" json:"product_id"`
	Quantity  int     `gorm:"not null;" json:"quantity"`
}
