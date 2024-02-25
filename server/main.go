package main

import (
	"net/http"

	"golang.org/x/net/websocket"

	"snowball-community.com/chat/db"
	"snowball-community.com/chat/features/message"
	"snowball-community.com/chat/features/room"
	"snowball-community.com/chat/features/user"
)

func main() {
	db.GetDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/users/username", user.GetUserName)

	// Rooms
	mux.HandleFunc("/rooms", room.Get)
	mux.Handle("/rooms/connect", websocket.Handler(room.HandleConnection))
	// mux.HandleFunc("/rooms/create", room.Create)

	//tMessages
	mux.HandleFunc("/messages", message.RouterHandler)
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		panic(err)
	}
}
