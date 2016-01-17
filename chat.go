package main

type Message struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Muted   bool   `json:"muted"`
}
