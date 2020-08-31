package main

import (
	"github.com/gorilla/websocket"
	"github.com/ha666/logs"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		logs.Error("r:%+v,stateus:%d,reason:%s", r, status, reason)
	},
}

func startWebsocket() {
	go statisticsClientTotal()
	http.HandleFunc("/echo", echo)
	logs.Emergency(http.ListenAndServe(addr, nil))
}
