package controllers

import (
	"encoding/json"
	"net/http"
	"snowball-community.com/chat/models"
	"snowball-community.com/chat/services"
	"snowball-community.com/chat/utils"
	"time"
)

func CreateChat(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only Post requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var chat models.Chat
	err := dec.Decode(&chat)
	if err != nil {
		panic(err)
	}

	chat.CreatedAt = time.Now().UnixMilli()
	chatId := services.CreateChat(chat)
	chat.ID = chatId

	utils.WriteEncoder(w, chat)
}
