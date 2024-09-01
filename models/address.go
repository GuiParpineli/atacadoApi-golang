package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Street       string `gorm:"size:255;not null;" json:"street"`
	Number       string `gorm:"size:255;not null;" json:"number"`
	Neighborhood string `gorm:"size:255;" json:"neighborhood"`
	City         string `gorm:"size:255;not null;" json:"city"`
	State        string `gorm:"size:255;not null;" json:"state"`
	ZipCode      string `gorm:"size:255;not null;" json:"zip_code"`
	County       string `gorm:"size:255;not null;" json:"county"`
	Complement   string `gorm:"size:255;" json:"complement"`
	CustomerID   uint   `json:"customer_id"`
}
