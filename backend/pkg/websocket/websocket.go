package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

func Reader(conn *websocket.Conn) {
	for {
		msgType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(p))
		if err := conn.WriteMessage(msgType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")
		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}
		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
