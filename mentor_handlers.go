package main

import (
	"net/http"

	"github.com/bigroom/communicator"
)

var mentors = []User{
	{
		Username:  "harrison",
		FirstName: "Harrison",
		LastName:  "Shoebridge",
		Email:     "harrison@theshoebridges.com",
		Bio:       "Half human, half hacker. AMA.",
	},
}

func getMentorsHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	coms.With(mentors).
		OK("Here are the mentors!")
}
