package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ha666/logs"
	"github.com/ha666/ws-server/service"
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
	go service.StatisticsClientTotal()
	http.HandleFunc("/process", process)
	logs.Emergency(http.ListenAndServe(addr, nil))
}
