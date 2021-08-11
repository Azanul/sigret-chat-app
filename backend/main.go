package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func reader(conn *websocket.Conn){
	for{
		msgType, p, err := conn.ReadMessage()
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(string(p))
		if err := conn.WriteMessage(msgType, p); err != nil{
			log.Println(err)
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func servews(w http.ResponseWriter, r* http.Request)  {
	fmt.Println(r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		fmt.Println(err)
	}
	reader(ws)
}
func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	http.HandleFunc("/ws", servews)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}