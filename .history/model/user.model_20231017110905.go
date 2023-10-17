package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time
}
