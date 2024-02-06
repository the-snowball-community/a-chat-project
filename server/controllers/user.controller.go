package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"snowball-community.com/chat/models"
	"snowball-community.com/chat/services"
)

func GetUserName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Expected Get request only", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-type", "application/json")
	now := time.Now()
	newUser := models.User{Name: services.CreateRandomUserName(), CreatedAt: now.UnixMilli()}
	user := services.CreateUser(newUser)

	err := json.NewEncoder(w).Encode(models.User{ID: user, Name: newUser.Name, CreatedAt: newUser.CreatedAt})
	if err != nil {
		panic(err)
	}
}
