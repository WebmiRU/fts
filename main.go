package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
)

var addr = flag.String("addr", "0.0.0.0:13900", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func accept(w http.ResponseWriter, r *http.Request) {
	c, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Print("UPGRADE:", e)
		return
	}

	defer c.Close()

	for {
		messageType, message, e := c.ReadMessage()

		if e != nil {
			log.Println("read:", e)
			break
		}

		log.Printf("REVICED:\n%s\n", message)

		e = c.WriteMessage(messageType, message)

		if e != nil {
			log.Println("write:", e)
			break
		}
	}
}

func main() {
	dsn := "host=192.168.1.151 user=postgres password=postgres dbname=fts port=5415 sslmode=disable TimeZone=Europe/Moscow"
	db, e := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
