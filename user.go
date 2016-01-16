package main

type Role int

const (
	MentorRole = iota + 1
	MenteeRole
)

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Bio        string `json:"bio"`
	Role       Role   `json:"role"`
	TalkingTo  []int  `json:"talking_to"`
	PictureURL string `json:"picture_url"`
}
