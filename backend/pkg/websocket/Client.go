package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	Key  []byte
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var newMsg Message
		err := c.Conn.ReadJSON(&newMsg)
		if err != nil {
			log.Println(err)
			return
		}

		if newMsg.Type == 0 {
			c.ChangeKey(newMsg.Body)
		} else if newMsg.Type == 1 {
			//fmt.Println("1: ", string(p))
			encrypted, _ := Encrypt([]byte(newMsg.Body), c.Key)
			//fmt.Println("2: ", string(p))
			c.SendMsg(encrypted)
		}
	}
}

func (c *Client) ChangeKey(newKey string) {
	c.Key = []byte(newKey)
}

func (c *Client) SendMsg(msg []byte) {
	message := Message{Type: 1, Body: string(msg)}
	c.Pool.Broadcast <- message
	//fmt.Printf("Message Received: %+v\n", message)
}
