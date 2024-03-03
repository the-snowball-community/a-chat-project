package room

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
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
	Ws      *websocket.Conn
	Channel string
}

var clients []Client

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ConnectRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	/// TODO: get a room number from db or user
	clients = append(clients, Client{Ws: conn, Channel: "65bace72b0200e4234de6e50"})
}

func GetClients() []Client {
	return clients
}

func HandleConnection(ws *websocket.Conn) {
	// fmt.Println("Trying Connection")
	// conn, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("Client connected")
	// conn.Close()
	// client := &Client{ws: ws}

	// clients[client] = true

	// log.Printf("Client connected")

	// for {
	// 	var msg map[string]interface{}
	// 	if err := websocket.JSON.Receive(ws, &msg); err != nil {
	// 		log.Println(err)
	// 		delete(clients, client)
	// 		break
	// 	}

	// 	action, ok := msg["action"].(string)
	// 	if !ok {
	// 		log.Println("Invalid action")
	// 		continue
	// 	}

	// 	switch action {
	// 	case "subscribe":
	// 		channel, ok := msg["channelId"].(string)
	// 		if !ok {
	// 			log.Println("Invalid channel ID")
	// 			continue
	// 		}
	// 		client.channel = channel
	// 		log.Printf("Client subscribed to channel %s", channel)

	// 	case "unsubscribe":
	// 		client.channel = ""
	// 		log.Printf("Client unsubscribed from channel")

	// 	default:
	// 		log.Println("Unknown action:", action)
	// 	}
	// }
}
