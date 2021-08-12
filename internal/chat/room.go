package chat

import (
	"github.com/google/uuid"
)

type Room struct {
	id uuid.UUID
	users        map[*User]bool
	Register     chan *User
	disconnected chan *User
	unregister chan *User
	broadcast chan *Message
}

func NewRoom() *Room {
	return &Room{
		id:           uuid.New(),
		users:        make(map[*User]bool),
		Register:     make(chan *User),
		disconnected: make(chan *User),
		unregister:   make(chan *User),
		broadcast:    make(chan *Message),
	}
}

func (room *Room) RunRoom ()  {
	for {
		select {
		case user := <- room.Register:
			room.registerUserInRoom(user)
		case user := <- room.disconnected:
			room.disconnectUserFromRoom(user)
		case user := <- room.unregister:
			room.removeFromRoom(user)
		case message := <- room.broadcast:
			room.broadcastMessageToRoom(message)
		}
	}
}

func (room *Room) registerUserInRoom (user *User) {
	room.users[user] = true
}

func (room *Room) removeFromRoom (user *User) {
	if _, exist := room.users[user]; exist {
		delete(room.users, user)
	}
}

func (room *Room) disconnectUserFromRoom (user *User) {
	if _, exist := room.users[user]; exist {
		room.users[user] = false
	}
}

func (room *Room) broadcastMessageToRoom(message *Message) {
	for user := range room.users {
		user.send <- message
	}
}