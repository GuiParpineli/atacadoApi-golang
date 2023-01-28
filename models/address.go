package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Street     string `gorm:"size:255;not null; unique" json:"street"`
	City       string `gorm:"size:255;not null; unique" json:"city"`
	State      string `gorm:"size:255;not null; unique" json:"state"`
	Zipcode    string `gorm:"size:255;not null; unique" json:"zipcode"`
	CustomerID uint
}
