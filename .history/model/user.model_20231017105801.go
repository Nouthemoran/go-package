package model

import "time"

type User struct {
	ID        int `gorm:"primaryKey" json:"id"`
	Name      string `json`
	CreatedAt time.Time
	UpdateAt  time.Time
}
