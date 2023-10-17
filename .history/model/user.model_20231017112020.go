package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    `json:"id"`
	Name string `json:"name"`
}
