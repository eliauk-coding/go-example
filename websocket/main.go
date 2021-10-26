package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"goUpUp/websocket/hub"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	go hub.H.Run()
	router.HandleFunc("/ws", hub.Task)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
