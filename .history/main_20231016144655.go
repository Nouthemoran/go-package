package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// CARA 1 
		// json.NewEncoder(w).Encode(map[string]bool{"ok": true})

		// cara 2 
		res, _ := json.Marshal(map[string]bool{"ok": true})
		
	}).Methods("GET")

	log.Println("Server running on port 9000")
	http.Handle("/", r) // Gunakan router mux sebagai handler
	http.ListenAndServe(":9000", nil)
}
