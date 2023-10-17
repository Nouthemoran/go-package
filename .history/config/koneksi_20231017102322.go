package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	connectionString := fmt.Sprin
	db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
