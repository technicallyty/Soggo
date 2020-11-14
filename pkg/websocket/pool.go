package websocket

import (
	"fmt"
	"log"
)

const (
	//ErrorJoiningRoom - error message sent to client when room not exist
	ErrorJoiningRoom = "Room does not exist."
)

//Pool - client pool
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Rooms      map[string][]*Client
	JoinRoom   chan *Client
	CreateRoom chan *Client
	Broadcast  chan Message
}

//NewPool cstrct
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Rooms:      make(map[string][]*Client),
		JoinRoom:   make(chan *Client),
		CreateRoom: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

//Start - begins the pool
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
		case client := <-pool.JoinRoom:
			fmt.Println("Client wants to join room.")
			_, ok := pool.Rooms[client.Room]
			if !ok {
				err := Message{Type: 0, Body: ErrorJoiningRoom}
				client.Conn.WriteJSON(err)
				log.Fatal(ErrorJoiningRoom)
			} else {
				pool.Rooms[client.Room] = append(pool.Rooms[client.Room], client)
			}
		}
	}
}
