package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

//Client struct
type Client struct {
	Nick string
	Conn *websocket.Conn
	Pool *Pool
	Room string
	mu   sync.Mutex
}

//Message struct
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
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		//	we receive our message here.
		//	we need to implement methods to handle
		//	different messages. for example
		//	commands and actions.

		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)

	}
}
