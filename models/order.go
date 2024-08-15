package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Products   []ProductOrder `gorm:"constraint:OnUpdate:CASCADE,constraint:OnDelete:CASCADE;"`
	Total      float32        `gorm:"not null;" json:"total"`
	Customer   Customer       `gorm:"constraint:OnUpdate:CASCADE,constraint:OnDelete:CASCADE;"`
	CustomerID uint           `json:"customer_id"`
	Vendor     Vendor         `gorm:"constraint:OnUpdate:CASCADE,constraint:OnDelete:CASCADE;"`
	VendorID   uint           `json:"vendor_id"`
}
