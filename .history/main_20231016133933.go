package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server")
	http.ListenAndServe(":9000", nil)
}
