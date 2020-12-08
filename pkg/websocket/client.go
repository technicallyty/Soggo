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

/*
*	Types -
*	1: User joined Room
*	2: User create room
* 3:
 */

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

		fmt.Printf("Message Type: %v\n Body: %v\n", messageType, string(p))

		//	we receive our message here.
		//	we need to implement methods to handle
		//	different messages. for example
		//	commands and actions.

		switch string(p) {
		case "1":
			fmt.Println("Joining room...")
			c.Room = "1234"
			c.Nick = "Tyler"
			c.Pool.JoinRoom <- c
		case "2":
			c.Room = "1234"
			c.Pool.CreateRoom <- c
		}
		/*
			message := Message{Type: messageType, Body: string(p)}
			c.Pool.Broadcast <- message
			fmt.Printf("Message Received: %+v\n", message)
		*/
	}
}
