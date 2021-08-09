package server

import (
	"github.com/SarapDev/chatWS/internal/chat"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WsServer struct {
	upgrade websocket.Upgrader
	room *chat.Room
	writer http.ResponseWriter
	request *http.Request
}

func NewWsServer (w http.ResponseWriter, r *http.Request, room *chat.Room) *WsServer {
	return &WsServer{
		upgrade: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		room: room,
		writer: w,
		request: r,
	}
}

func (ws *WsServer) Serve (room *chat.Room) {
	conn, err := ws.upgrade.Upgrade(ws.writer, ws.request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user := chat.NewUser(conn, room)
	user.Room.Register <- user

	go user.ReadMessage()
	go user.SendMessage()
}