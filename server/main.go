package main

import (
	"log"
	"net/http"
	"time"

	"snowball-community.com/chat/db"
)

func main() {
	db.GetDB()
	app := application{}

	svr := &http.Server{
		Addr:         ":4000",
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := svr.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
