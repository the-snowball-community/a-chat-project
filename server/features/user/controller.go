package user

import (
	"encoding/json"
	"net/http"
	"time"

	"snowball-community.com/chat/models"
)

func GetUserName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodGet {
		http.Error(w, "Expected Get request only", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-type", "application/json")
	now := time.Now()
	newUser := models.User{Name: CreateRandomUserName(), CreatedAt: now.UnixMilli()}
	user := CreateUser(newUser)

	err := json.NewEncoder(w).Encode(models.User{ID: user, Name: newUser.Name, CreatedAt: newUser.CreatedAt})
	if err != nil {
		panic(err)
	}
}
