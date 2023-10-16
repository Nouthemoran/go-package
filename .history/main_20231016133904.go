package main

import (
	"net/http"
)

func main() {
	log
	http.ListenAndServe(":9000", nil)
}