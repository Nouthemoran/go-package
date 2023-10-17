package model

import "time"

type User struct {
	ID      string `json:"id"`
	Name    string
	Created time.Time
	Updated time.Time
}
