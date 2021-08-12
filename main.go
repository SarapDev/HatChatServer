package main

import (
	"flag"
	"github.com/SarapDev/chatWS/internal/chat"
	server2 "github.com/SarapDev/chatWS/internal/server"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request)  {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main()  {
	flag.Parse()
	room := chat.NewRoom()
	go room.RunRoom()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server := server2.NewWsServer(w, r, room)
		server.Serve()
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}