package main

import (
	"chat/websock"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var address = flag.String("address", ":8081", "HTTP serve")

func main() {
	flag.Parse()
	hub := websock.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websock.ServeWs(hub, w, r)
	})

	fs := http.FileServer(http.Dir("../frontend/Chat/dist"))
	http.Handle("/", fs)

	fmt.Println("Chat server opened")

	log.Fatal(http.ListenAndServe(*address, nil))
}
