package model

import "time"

type User struct {
	ID        int 
	Name      string 
	CreatedAt time.Time 
	UpdateAt  time.Time
}
