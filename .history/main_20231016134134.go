package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux

	log.Println("Server running on port 9000")
	http.ListenAndServe(":9000", nil)
}
