package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var (
	users map[string]*websocket.Conn
)

func main() {
	users = make(map[string]*websocket.Conn)

	r := mux.NewRouter()

	r.HandleFunc("/mentees", getMenteesHandler).Methods("GET")
	r.HandleFunc("/mentees/{id}", getMenteeHandler).Methods("GET")

	r.HandleFunc("/mentors", getMentorsHandler).Methods("GET")
	r.HandleFunc("/mentors/{id}", getMentorHandler).Methods("GET")

	r.HandleFunc("/chat", chatHandler)
	r.HandleFunc("/chat/messages", getChatHandler).Methods("GET")
	r.HandleFunc("/chat/messages", postMessageHandler).Methods("POST")

	http.Handle("/", cors.Default().Handler(r))

	log.Println("Serving on 0.0.0.0:8080!")

	http.ListenAndServe("0.0.0.0:8080", nil)
}
