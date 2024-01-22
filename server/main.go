package main

import (
	"net/http"

	"snowball-community.com/chat/controllers"
	"snowball-community.com/chat/db"
)

func main() {
	db.GetDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/username", controllers.GetUserName)

	// Rooms
	mux.HandleFunc("/room", controllers.GetRoom)
	//mux.HandleFunc("/room/create", controllers.CreateRoom)
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		panic(err)
	}
}
