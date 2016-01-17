package main

import (
	"net/http"
	"strconv"

	"github.com/bigroom/communicator"
	"github.com/gorilla/mux"
)

var mentees = []User{
	{
		ID:         5,
		Username:   "rafej",
		FirstName:  "Rafe",
		LastName:   "Skidmore",
		Email:      "hello@rafej.io",
		Bio:        "Making the world a better place by making people smile!",
		Role:       MenteeRole,
		PictureURL: "https://avatars1.githubusercontent.com/u/13427998?v=3&s=460",
	},
	{
		ID:         6,
		Username:   "harrison",
		FirstName:  "Harrison",
		LastName:   "Shoebridge",
		Email:      "harrison@theshoebridges.com",
		Bio:        "Half human, half hacker.",
		Role:       MenteeRole,
		PictureURL: "https://avatars1.githubusercontent.com/u/8407370?v=3&s=460",
		TalkingTo: []int{
			0, // Alex Hogue
			1, // Zach Latta
			3, // Peter Thiel
		},
	},
	{
		ID:         7,
		Username:   "tcurran",
		FirstName:  "Thomas",
		LastName:   "Curran",
		Email:      "thomas@curran.com",
		Bio:        "Son of James Curran, god of Computational Linguistics.",
		Role:       MenteeRole,
		PictureURL: "http://www.todaysparent.com/wp-content/uploads/2011/09/ProfileGrowing-Baby.jpg",
	},
}

func getMenteesHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	coms.With(mentees).
		OK("Here are your mentees!")
}

func getMenteeHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	i, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		coms.With(err).
			Fail("Could no cast id")
		return
	}

	coms.With(mentees[i-len(mentors)]).
		OK()
}
