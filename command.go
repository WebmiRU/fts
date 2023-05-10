package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

func checkUserLoggedIn(socket *websocket.Conn) bool {
	if _, ok := sockets[socket]; !ok {
		var response, _ = json.Marshal(MessageAuthLoginError) // @TODO Change reponse
		socket.WriteMessage(websocket.TextMessage, response)
		socket.Close()

		return false
	}

	return true
}

func cmdAuthLogin(token string) User {
	var model User
	db.Preload("Channels").Where("token = ?", token).First(&model)

	return model
}

func cmdChannelList(userId uint64) []Channel {
	var model User
	db.Preload("Channels").First(&model, userId)

	return model.Channels
}
