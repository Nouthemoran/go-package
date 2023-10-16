package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("S")
	http.ListenAndServe(":9000", nil)
}
