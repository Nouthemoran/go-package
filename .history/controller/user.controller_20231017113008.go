package controller

import (
	"go-package/config"
	"go-package/model"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		w.
	}


}
