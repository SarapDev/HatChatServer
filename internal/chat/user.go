package chat

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type User struct {
	id uuid.UUID
	conn *websocket.Conn
	send chan *Message
	Room *Room
}

func NewUser(conn *websocket.Conn, room *Room) *User {
	return &User{
		id:   uuid.New(),
		conn: conn,
		send: make(chan *Message),
		Room: room,
	}
}

func (user *User) SendMessage () {

}

func (user *User) ReadMessage () {

}