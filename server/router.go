package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/websocket"
	"snowball-community.com/chat/features/message"
	"snowball-community.com/chat/features/room"
	"snowball-community.com/chat/features/user"
)

type application struct{}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"status": "available",
		"port":   ":4000",
	}
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	/// Health Check
	mux.HandleFunc("/health", healthcheck)

	/// Users
	mux.HandleFunc("/users/username", user.GetUserName)

	/// Rooms
	mux.HandleFunc("/rooms", room.Get)
	mux.Handle("/rooms/connect", websocket.Handler(room.HandleConnection))

	/// Messages
	mux.HandleFunc("/messages", message.RouterHandler)

	return mux
}
