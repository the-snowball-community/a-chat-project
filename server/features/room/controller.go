package room

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	"snowball-community.com/chat/utils"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	now := time.Now()
	newRoom := Room{CreatedAt: now.UnixMilli()}
	user := CreateRoom(newRoom)

	utils.WriteEncoder(w, Room{ID: user, CreatedAt: newRoom.CreatedAt})
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	utils.WriteEncoder(w, GetRoom())
}

type Client struct {
	ws      *websocket.Conn
	channel string
}

var clients = make(map[*Client]bool)

func HandleConnection(ws *websocket.Conn) {
	client := &Client{ws: ws}

	clients[client] = true

	log.Printf("Client connected")

	for {
		var msg map[string]interface{}
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			log.Println(err)
			delete(clients, client)
			break
		}

		action, ok := msg["action"].(string)
		if !ok {
			log.Println("Invalid action")
			continue
		}

		switch action {
		case "subscribe":
			channel, ok := msg["channelId"].(string)
			if !ok {
				log.Println("Invalid channel ID")
				continue
			}
			client.channel = channel
			log.Printf("Client subscribed to channel %s", channel)

		case "unsubscribe":
			client.channel = ""
			log.Printf("Client unsubscribed from channel")

		default:
			log.Println("Unknown action:", action)
		}
	}
}
