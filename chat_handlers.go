package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	user, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	defer user.Close()

	username := r.FormValue("username")

	users[username] = user

	log.Info("New user connected! " + username)

	for {
		var m Message

		err := user.ReadJSON(&m)
		if err != nil {
			log.Warn("Closing connection. Error reading:", err)

			return
		}

		to := users[m.Username]
		if to == nil {
			log.Println("User does not exist", m.Username)
			continue
		}

		to.WriteJSON(Message{
			Message:  m.Message,
			Username: username,
			Muted:    m.Muted,
		})
	}
}
