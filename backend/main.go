package main

import (
	"fmt"
	"github.com/Azanul/sigret-chat-app/pkg/websocket"
	"net/http"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Println(err)
	}

	client := &websocket.Client{
		Conn: ws,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
