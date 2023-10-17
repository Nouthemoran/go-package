package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Mod
	ID      string `json:"id"`
	Name    string
	Created time.Time
	Updated time.Time
}
