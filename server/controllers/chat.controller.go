package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"snowball-community.com/chat/models"
	"snowball-community.com/chat/services"
	"snowball-community.com/chat/utils"
)

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method == http.MethodGet {
		loadChats(w, r)
	} else if r.Method == http.MethodPost {
		createChat(w, r)
	} else {
		http.Error(w, "This request method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func createChat(w http.ResponseWriter, r *http.Request) {
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

func loadChats(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil || limit < 0 {
		http.NotFound(w, r)
	}
	skip, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)
	if err != nil || skip < 0 {
		http.NotFound(w, r)
	}
	var chats []models.Chat = services.FindMany(limit, skip)

	utils.WriteEncoder(w, chats)
}
