package models

type Address struct {
	Street string `gorm:"size:255;not null; unique" json:"street"`
}
