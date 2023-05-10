package main

type MessagePayload struct {
	Token     string `json:"token"`
	ChannelId uint64 `json:"channel_id"`
	Message   string
	Text      string
}

type Message struct {
	Type    string         `json:"type"`
	Payload MessagePayload `json:"payload"`
}
