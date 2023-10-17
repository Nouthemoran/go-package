package model

import "time"

type User struct {
	ID        int `gorm`
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
}
