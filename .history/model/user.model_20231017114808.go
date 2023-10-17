package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
