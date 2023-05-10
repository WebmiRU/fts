package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
	"strings"
)

var db *gorm.DB
var sockets = make(map[*websocket.Conn]*User)
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

		if err != nil {
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
			var user = cmdAuthLogin(msg.Payload.Token)

			if user.ID == 0 {
				var response, _ = json.Marshal(MessageAuthLoginError)
				socket.WriteMessage(websocket.TextMessage, response)
				socket.Close() // Auth error
				break
			}

			sockets[socket] = &user

			var response, _ = json.Marshal(MessageAuthLoginOK)
			socket.WriteMessage(websocket.TextMessage, response)

			println("ID:", user.ID)
			fmt.Printf("\nCHANNELS:\n%+v\n:", sockets)
			break

		case "channel/list":
			if !checkUserLoggedIn(socket) {
				break
			}

			var channels = messageUserChannelList{Payload: sockets[socket].Channels}
			var response, _ = json.Marshal(channels)

			socket.WriteMessage(websocket.TextMessage, response)

			break
		}

		log.Printf("REVICED:\n%s\n", msg)
	}
}

func main() {
	var dsn = "host=localhost user=fts password=fts dbname=fts port=5432 sslmode=disable TimeZone=Europe/Moscow"
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
