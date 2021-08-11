package main

import (
	"fmt"
	"github.com/Azanul/sigret-chat-app/pkg/websocket"
	"net/http"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Println(err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}
func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	http.HandleFunc("/ws", serveWs)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
