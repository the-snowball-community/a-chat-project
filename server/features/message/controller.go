package message

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"snowball-community.com/chat/utils"
)

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin, Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Content-type", "application/json")
	if r.Method == http.MethodGet {
		load(w, r)
	} else if r.Method == http.MethodPost {
		create(w, r)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		http.Error(w, "This request method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	// var message Message
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }

	// err = json.Unmarshal(body, &message)
	// if err != nil {
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var message Message
	err := dec.Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	message.CreatedAt = time.Now().UnixMilli()
	messageId := CreateMessage(message)
	message.ID = messageId

	utils.WriteEncoder(w, message)
}

func load(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil || limit < 0 {
		http.NotFound(w, r)
	}
	skip, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)
	if err != nil || skip < 0 {
		http.NotFound(w, r)
	}
	var messages []Message = FindMany(limit, skip)

	utils.WriteEncoder(w, messages)
}
