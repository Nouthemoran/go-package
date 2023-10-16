package main

import (
	"net/http"
)

func main() {
	log.Print
	http.ListenAndServe(":9000", nil)
}
