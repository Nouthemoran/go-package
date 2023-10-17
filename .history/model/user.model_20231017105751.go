package model

import "time"

type User struct {
	ID        int `gorm:"primaryKey" json:"id"`
	Name      string ~
	CreatedAt time.Time
	UpdateAt  time.Time
}
