package main

import (
	"net/http"

	"golang.org/x/net/websocket"

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
	mux.Handle("/rooms/connect", websocket.Handler(controllers.ConnectRoom))

	//tMessages
	mux.HandleFunc("/messages", controllers.RouterHandler)
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		panic(err)
	}
}
