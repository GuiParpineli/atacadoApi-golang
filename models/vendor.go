package models

import "github.com/jinzhu/gorm"

type Vendor struct {
	gorm.Model
	Name     string  `gorm:"size:255; not null;" json:"name"`
	Lastname string  `gorm:"size:255; not null;" json:"lastname"`
	Cpf      string  `gorm:"size:255; not null;" json:"cpf"`
	Email    string  `gorm:"size:255; not null;" json:"email"`
	Phone    string  `gorm:"size:255;not null; unique" json:"phone"`
	Address  Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
