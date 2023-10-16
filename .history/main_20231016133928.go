package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println()
	http.ListenAndServe(":9000", nil)
}
