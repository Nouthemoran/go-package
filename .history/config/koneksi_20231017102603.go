package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connectionString := fmt.Sprintf("%v:%v@/%v?parseTime=true". ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_DATABASE)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
