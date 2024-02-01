package main

import (
	"net/http"

	"snowball-community.com/chat/controllers"
	"snowball-community.com/chat/db"
)

func main() {
	db.GetDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/users/username", controllers.GetUserName)

	// Rooms
	mux.HandleFunc("/rooms", controllers.GetRoom)
	// mux.HandleFunc("/room/create", controllers.CreateRoom)

	// Chat
	mux.HandleFunc("/messages", controllers.RouterHandler)
	// mux.HandleFunc("/chat/send", controllers.CreateChat)
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		panic(err)
	}
}
