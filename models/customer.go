package models

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Name     string  `gorm:"size:255;not null;" json:"name"`
	Lastname string  `gorm:"size:255;not null;" json:"lastname"`
	Email    string  `gorm:"size:255;not null; unique" json:"email"`
	Phone    string  `gorm:"size:255;not null; unique" json:"phone"`
	Cpnj     string  `gorm:"size:255;not null; unique" json:"cpnj"`
	Address  Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
