package chat

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
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

func (u *User) SendMessage () {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := u.conn.Close()

		if err != nil {
			return
		}
	}()

	for {
		select {
		case message, ok := <-u.send:
			err := u.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if err != nil {
				return
			}

			if !ok {
				u.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := u.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, err = w.Write(message.Text)
			if err != nil {
				return
			}

			n := len(u.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				text := <-u.send
				w.Write(text.Text)

				log.Println("User: ", u.id)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (u *User) ReadMessage () {
	defer func () {
		u.Room.unregister <- u
		err := u.conn.Close()
		if err != nil {
			return
		}
	} ()

	u.conn.SetReadLimit(maxMessageSize)

	err := u.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return
	}

	u.conn.SetPongHandler(func (string) error {
		err := u.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			return err
		}
		return nil
	})

	for {
		_, message, err := u.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		messageStruct := NewMessage(message, u)
		u.Room.broadcast <- messageStruct
	}
}