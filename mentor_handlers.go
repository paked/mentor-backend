package main

import (
	"net/http"
	"strconv"

	"github.com/bigroom/communicator"
	"github.com/gorilla/mux"
)

var mentors = []User{
	{
		ID:        0,
		Username:  "harrison",
		FirstName: "Harrison",
		LastName:  "Shoebridge",
		Email:     "harrison@theshoebridges.com",
		Bio:       "Half human, half hacker. AMA.",
		Role:      MentorRole,
	},
	{
		ID:        1,
		Username:  "zrl",
		FirstName: "Zach",
		LastName:  "Latta",
		Email:     "zach@hackclub.io",
		Bio:       "Some hybrid of something or other?",
		Role:      MentorRole,
	},
	{
		ID:        2,
		Username:  "aakagi",
		FirstName: "Alex",
		LastName:  "Akagi",
		Email:     "alex@alexakagi.something",
		Bio:       "Startups and shit.",
		Role:      MentorRole,
	},
	{
		ID:        3,
		Username:  "pthiel",
		FirstName: "Peter",
		LastName:  "Thiel",
		Email:     "me@peterthiel.com",
		Bio:       "AMA.",
		Role:      MentorRole,
		TalkingTo: []int{
			5, // Rafe Skidmore
			6, // Alex Hogue
			7, // Thomas Curran
		},
	},
	{
		ID:        4,
		Username:  "kim.com",
		FirstName: "Kim",
		LastName:  "Dotcom",
		Email:     "me@kim.com",
		Bio:       "I own the cyber space?",
		Role:      MentorRole,
	},
}

func getMentorsHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	coms.With(mentors).
		OK("Here are the mentors!")
}

func getMentorHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	i, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		coms.With(err).
			Fail("Could not convert string")
		return
	}

	coms.With(mentors[i]).
		OK()
}
