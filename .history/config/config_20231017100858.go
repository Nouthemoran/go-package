package config

import "github.com/spf13/viper"

type config struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
}

var ENV *config

func LoadConfig() {
	viper.AddConfigPath()
}
