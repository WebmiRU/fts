package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/slices"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
	"strings"
)

var db *gorm.DB
var sockets = make(map[*websocket.Conn]*User)
var channels = make(map[uint64][]*websocket.Conn)
var addr = flag.String("addr", "0.0.0.0:13900", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func accept(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("UPGRADE:", err)
		return
	}

	defer socket.Close()

	for {
		_, message, err := socket.ReadMessage()

		if err != nil { // Connection closed
			delete(sockets, socket)

			for channelId := range channels {
				//if _, ok := sockets[socket]; !ok {
				if idx := slices.Index(channels[channelId], socket); idx >= 0 {
					channels[channelId] = slices.Delete(channels[channelId], idx, idx)
					fmt.Println("DELETE INDEX:", idx)
				}

				//if idx >= 0 { // Delete socket from channel list
				//	channels[channelId] = slices.Delete(channels[channelId], idx, idx)
				//	//channels[channelId][idx] = channels[channelId][len(channels[channelId])-1]
				//	//channels[channelId] = channels[channelId][:len(channels[channelId])-1]
				//}
			}

			log.Println("read:", err)
			break
		}

		var msg Message
		err = json.Unmarshal(message, &msg)

		if err != nil {
			fmt.Println("JSON PARSE ERROR", err)
		}

		switch strings.ToLower(msg.Type) {
		case "auth/login":
			user := cmdAuthLogin(msg.Payload.Token)

			if user.ID == 0 {
				response, _ := json.Marshal(MessageAuthLoginError)
				socket.WriteMessage(websocket.TextMessage, response)
				socket.Close() // Auth error
				break
			}

			sockets[socket] = &user

			for _, v := range user.Channels {
				channels[v.ID] = append(channels[v.ID], socket)
				fmt.Printf("\n\n%+v\n\n", channels)
			}

			response, _ := json.Marshal(MessageAuthLoginOK)
			socket.WriteMessage(websocket.TextMessage, response)

			println("ID:", user.ID)
			fmt.Printf("\nCHANNELS:\n%+v\n:", sockets)
			break

		case "channel/list":
			if !userLoggedIn(socket) {
				break
			}

			channels := messageUserChannelList{Payload: sockets[socket].Channels}
			response, _ := json.Marshal(channels)

			socket.WriteMessage(websocket.TextMessage, response)

			break

		case "channel/message":
			if !userLoggedIn(socket) {
				break
			}

			channelMessage := ChannelMessage{
				UserId:    sockets[socket].ID,
				ChannelId: msg.Payload.ChannelId,
				Message:   msg.Payload.Message,
				Text:      msg.Payload.Text,
			}

			m, _ := json.Marshal(channelMessage)

			// Send channel clients message
			// @TODO Add check user in channel before send message
			for _, s := range channels[msg.Payload.ChannelId] {
				s.WriteMessage(websocket.TextMessage, m)
			}

			break
		}

		log.Printf("REVICED:\n%s\n", msg)
	}
}

func main() {
	dsn := "host=192.168.1.151 user=fts password=fts dbname=fts port=5415 sslmode=disable TimeZone=Europe/Moscow"
	var e error

	db, e = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	var user []User

	db.Preload("Channels").Find(&user)

	for _, v := range user {
		fmt.Printf("%+v\n", v.Channels)
	}

	//fmt.Printf("%+v\n", user)
	fmt.Println(db, e)

	http.HandleFunc("/", accept)
	log.Fatal(http.ListenAndServe(*addr, nil))

	fmt.Println("Hello world")
}
