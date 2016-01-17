package main

import (
	"net/http"
	"strconv"

	"github.com/bigroom/communicator"
	"github.com/gorilla/mux"
)

var mentors = []User{
	{
		ID:         0,
		Username:   "scrub",
		FirstName:  "Alex",
		LastName:   "Hogue",
		Email:      "alex@alexhogue.com",
		Bio:        "I work at Atlassian and am known as the one true gagmeister.",
		Role:       MentorRole,
		PictureURL: "http://img.uefa.com/imgml/TP/players/3/2016/324x324/95417.jpg",
	},
	{
		ID:         1,
		Username:   "zrl",
		FirstName:  "Zach",
		LastName:   "Latta",
		Email:      "zach@hackclub.io",
		Bio:        "Some hybrid of something or other? I like EDtech, and bash.",
		Role:       MentorRole,
		PictureURL: "https://avatars1.githubusercontent.com/u/992248?v=3&s=460",
	},
	{
		ID:         2,
		Username:   "aakagi",
		FirstName:  "Alex",
		LastName:   "Akagi",
		Email:      "alex@alexakagi.something",
		Bio:        "Startups and shit. Also: Purdue",
		Role:       MentorRole,
		PictureURL: "https://pbs.twimg.com/profile_images/580237836739186688/_edKynPi.jpg",
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
		PictureURL: "http://static4.bornrichimages.com/cdn2/500/500/91/c/wp-content/uploads/2014/09/ptr_thumb.jpg",
	},
	{
		ID:         4,
		Username:   "kim.com",
		FirstName:  "Kim",
		LastName:   "Dotcom",
		Email:      "me@kim.com",
		Bio:        "You own the future. You own the world.",
		Role:       MentorRole,
		PictureURL: "http://www.wired.com/images_blogs/threatlevel/2012/01/Kim-Dotcom-with-Mercedes-Reuters.jpg",
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
