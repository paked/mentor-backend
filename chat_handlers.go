package main

import (
	"net/http"
	"strconv"

	"github.com/bigroom/communicator"
)

var messages []Message

func getChatHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	coms.With(messages).
		OK()
}

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)
	id, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		coms.Fail("shit happened")
	}

	m := Message{
		ID:      id,
		Message: r.FormValue("content"),
		Muted:   false,
	}

	messages = append(messages, m)

	coms.With(m).
		OK()
}
