package main

import (
	"net/http"

	"github.com/bigroom/communicator"
)

var mentees = []User{
	{
		Username:  "rafej",
		FirstName: "Rafe",
		LastName:  "Skidmore",
	},
}

func getMenteesHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	coms.With(mentees).
		OK("Here are your mentees!")
}
