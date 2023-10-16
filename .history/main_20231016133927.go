package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println(v)
	http.ListenAndServe(":9000", nil)
}
