package main

import (
	"net/http"

	"github.com/bigroom/communicator"
)

var mentees = []User{
	{
		ID:        5,
		Username:  "rafej",
		FirstName: "Rafe",
		LastName:  "Skidmore",
		Email:     "hello@rafej.io",
		Bio:       "Doing things. Doing them now.",
		Role:      MenteeRole,
	},
	{
		ID:        6,
		Username:  "scrub",
		FirstName: "Alex",
		LastName:  "Hogue",
		Email:     "alex@alexhogue.com",
		Bio:       "Doing work. Yeah, fun.",
		Role:      MenteeRole,
	},
	{
		ID:        7,
		Username:  "tcurran",
		FirstName: "Thomas",
		LastName:  "Curran",
		Email:     "thomas@curran.com",
		Bio:       "I am jesus. OK?",
		Role:      MenteeRole,
	},
}

func getMenteesHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	coms.With(mentees).
		OK("Here are your mentees!")
}

/*
func getMentorByID(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	coms.With(mentees[]
}*/
