package controller

import (
	"encoding/json"
	"go-package/config"
	"go-package/model"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		w.WriteHeader(500)
		res, _ := json.Marshal(map[string]string{"status": "failed"})
		w.Write(res)
	}

	res, _ := json.Marshal()

}
