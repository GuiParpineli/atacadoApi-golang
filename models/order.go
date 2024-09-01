package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Products        []ProductOrder `gorm:"constraint:OnUpdate:CASCADE,constraint:OnDelete:CASCADE;"`
	Status          string         `gorm:"not null;" json:"status"`
	PaymentMehod    string         `gorm:"not null;" json:"payment_method"`
	ShippingMethod  string         `gorm:"not null;" json:"shipping_method"`
	ShippingCost    float32        `gorm:"not null;" json:"shipping_cost"`
	ShippingAddress string         `gorm:"not null;" json:"shipping_address"`
	BillingAddress  string         `gorm:"not null;" json:"billing_address"`
	Customer        Customer       `gorm:"constraint:OnUpdate:CASCADE,constraint:OnDelete:CASCADE;"`
	CustomerID      uint           `json:"customer_id"`
	Vendor          Vendor         `gorm:"constraint:OnUpdate:CASCADE,constraint:OnDelete:CASCADE;"`
	VendorID        uint           `json:"vendor_id"`
	Total           float32        `gorm:"not null;" json:"total"`
}
