package websocket

import (
	"fmt"
)

//Room - data from room
type Room struct {
	RoomID    string
	Clients   map[*Client]bool
	Host      *Client
	Execute   chan Command
	GameState *GameState
}

//Command - the formula to issue commands to a room.
type Command struct {
	Sender  *Client
	Command string
	Type    int
}

//NewRoom - constructor for new rooms
func NewRoom(id string, host *Client) *Room {

	clientMap := make(map[*Client]bool)
	clientMap[host] = true

	room := &Room{
		RoomID:    id,
		Clients:   clientMap,
		Host:      host,
		Execute:   make(chan Command),
		GameState: &GameState{make(map[*Client]*UserState)},
	}

	//	create room in new goroutine as to not block main
	go room.start()

	return room
}

//May or may not to be exported, need to check that out.
func (room *Room) start() {
	for {
		select {
		case command := <-room.Execute:
			fmt.Println(command)
		}
	}
}

//BroadcastJoin - broadcasts to the clients a new members has joined
func (room *Room) BroadcastJoin(name string) {
	msg := Message{Type: 1, Body: string(name)}
	for client := range room.Clients {
		client.Conn.WriteJSON(msg)
	}
}
