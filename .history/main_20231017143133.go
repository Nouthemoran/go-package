package main

import (
	"encoding/json"
	"go-package/config"
	"go-package/controller"
	"go-package/middleware"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/users", controller.Index).Methods("GET")

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// CARA 1
		// json.NewEncoder(w).Encode(map[string]bool{"ok": true})

		// cara 2
		res, _ := json.Marshal(map[string]bool{"ok": true})
		w.Write(res)
	}).Methods("GET")

	r.Use(middleware.LoggingMiddleware)

	log.Println("Server running on port 9000")
	http.Handle("/", r) // Gunakan router mux sebagai handler
	http.ListenAndServe(":9000", nil)
} 
