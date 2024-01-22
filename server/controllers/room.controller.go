package controllers

import (
	"net/http"
	"snowball-community.com/chat/models"
	"snowball-community.com/chat/services"
	"snowball-community.com/chat/utils"
	"time"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	now := time.Now()
	newRoom := models.Room{CreatedAt: now.UnixMilli()}
	user := services.CreateRoom(newRoom)

	utils.WriteEncoder(w, models.Room{ID: user, CreatedAt: newRoom.CreatedAt})
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	utils.WriteEncoder(w, services.GetRoom())

}
