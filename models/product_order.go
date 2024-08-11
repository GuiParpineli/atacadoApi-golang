package models

import "github.com/jinzhu/gorm"

type ProductOrder struct {
	gorm.Model
	Product  Product `gorm:"constraint:OnUpdate:CASCADE,constraint:OnDelete:CASCADE;"`
	Quantity int     `gorm:"not null;" json:"quantity"`
}
