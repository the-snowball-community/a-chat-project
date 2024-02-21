package controllers

import (
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	"snowball-community.com/chat/models"
	"snowball-community.com/chat/services"
	"snowball-community.com/chat/utils"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	now := time.Now()
	newRoom := models.Room{CreatedAt: now.UnixMilli()}
	user := services.CreateRoom(newRoom)

	utils.WriteEncoder(w, models.Room{ID: user, CreatedAt: newRoom.CreatedAt})
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	utils.WriteEncoder(w, services.GetRoom())
}

func ConnectRoom(ws *websocket.Conn) {
	// websocket.Message.Receive(ws)
}
