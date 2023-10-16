package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server running on port ")
	http.ListenAndServe(":9000", nil)
}
