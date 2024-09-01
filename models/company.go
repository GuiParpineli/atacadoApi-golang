package models

import "github.com/jinzhu/gorm"

type Company struct {
	gorm.Model
	Cnpj        string   `gorm:"size:13;unique;not null"`
	CompanyName string   `gorm:"size:255;not null"`
	CompanyType string   `gorm:"size:255;not null"`
	Address     Address  `gorm:"size:255;not null"`
	AddressId   uint     `json:"address_id"`
	Customer    Customer `json:"customer"`
	CustomerId  uint     `json:"customer_id"`
}
