package config

func ConnectDB() {
	func main() {
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
		  panic
}