package server

import (
	"github.com/SarapDev/chatWS/internal/chat"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WsServer struct {
	upgrade websocket.Upgrader
	room    *chat.Room
	writer  http.ResponseWriter
	request *http.Request
}

func NewWsServer(w http.ResponseWriter, r *http.Request, room *chat.Room) *WsServer {
	return &WsServer{
		upgrade: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		room:    room,
		writer:  w,
		request: r,
	}
}

func (ws *WsServer) Serve() {
	conn, err := ws.upgrade.Upgrade(ws.writer, ws.request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	_, usernameMsg, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, usernameMsg)
	if err != nil {
		log.Println("ERROR, Ошибка отправки данных о подключенном пользователе")
	}

	user := chat.NewUser(conn, ws.room, string(usernameMsg))
	user.Room.Register <- user

	log.Println("Run Server")
	log.Println(user.Nickname)

	go user.ReadMessage()
	go user.SendMessage()
}
