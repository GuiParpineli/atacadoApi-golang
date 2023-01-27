package model

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name     string `gorm:"size:255;not null; unique" json:"name"`
	Lastname string `gorm:"size:255;not null; unique" json:"lastname"`
	Email    string `gorm:"size:255;not null; unique" json:"email"`
	Phone    string `gorm:"size:255;not null; unique" json:"phone"`
	Cpf      string `gorm:"size:255;not null; unique" json:"cpf"`
}
