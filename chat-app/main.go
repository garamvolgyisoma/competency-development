package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ //top-level variables must be defined with var keyword, in other places, it can be the := syntax
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println(http.ListenAndServe(":6969", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) { // *-syntax: represents a pointer
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(string(msg))
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}
