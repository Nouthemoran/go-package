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
	log
	viper.AddConfigPath('.')
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&ENV)
	if err != nil {
		panic(err)
	}
}
