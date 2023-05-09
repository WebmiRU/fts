package main

type MessagePayload struct {
	Token string `json:"token"`
}

type Message struct {
	Type    string         `json:"type"`
	Payload MessagePayload `json:"payload"`
}
