package models

import "github.com/jinzhu/gorm"

type ProductImage struct {
	gorm.Model
	Url string `json:"url"`
}
