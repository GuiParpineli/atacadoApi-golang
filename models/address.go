package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Street     string `gorm:"size:255;not null;" json:"street"`
	City       string `gorm:"size:255;not null;" json:"city"`
	State      string `gorm:"size:255;not null;" json:"state"`
	Zipcode    string `gorm:"size:255;not null;" json:"zipcode"`
	CustomerID uint
}
