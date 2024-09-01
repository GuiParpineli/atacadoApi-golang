package models

import "github.com/jinzhu/gorm"

type UserApp struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);unique_index"`
	Password string `gorm:"type:varchar(255),not null" json:"password"`
	Email    string `gorm:"type:varchar(255);unique_index" json:"email"`
	Phone    string `gorm:"type:varchar(255);unique_index" json:"phone"`
}
