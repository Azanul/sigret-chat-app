package main

import (
	"fmt"
	"github.com/Azanul/sigret-chat-app/pkg/websocket"
	"log"
	"net/http"
	"os"
)
//
//func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
//	w.WriteHeader(status)
//	if status == http.StatusNotFound {
//		log.Fatal(fmt.Fprint(w, "custom 404"))
//	}
//}
//
//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		errorHandler(w, r, http.StatusNotFound)
//		return
//	}
//	fmt.Println(r.Host)
//
//	http.FileServer(http.Dir("./backend/build"))
//}

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
	//http.HandleFunc("/", homeHandler)
	http.Handle("/", http.FileServer(http.Dir("./build")))
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	if r.URL.Path == "/" {
	//		http.FileServer(http.Dir("./backend/build"))
	//	} else {
	//		log.Fatal(fmt.Fprintln(w, "custom 404"))
	//	}
	//})

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	setupRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
