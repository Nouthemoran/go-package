package main

import (
	"log"
	"net/http"
)

func main() {
	r

	log.Println("Server running on port 9000")
	http.ListenAndServe(":9000", nil)
}
