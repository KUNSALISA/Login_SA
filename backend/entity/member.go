package entity

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
    Username string `json:"user_name"`
    Email     string `json:"email" gorm:"unique"`
    Password  string `json:"-"`
}