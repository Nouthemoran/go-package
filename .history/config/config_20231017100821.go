package config

type config struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
}

var ENV *config

func LoadConfig() {

}
