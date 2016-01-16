package main

import (
	"net/http"

	"github.com/bigroom/communicator"
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

func addMentorsForMenteeHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)
}
