package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	connectionString
	db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
