package main

import (
	"net/http"
	"time"

	"github.com/ha666/golibs"
	"github.com/ha666/logs"
	"github.com/ha666/ws-common"
	"github.com/ha666/ws-common/protocol"
)

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Error("upgrade:", err)
		return
	}
	clientAddr := ClientIpPort(r)
	if golibs.Length(clientAddr) <= 0 {
		if err = c.Close(); err != nil {
			logs.Error("当前连接没有ip，自动断开，断开连接失败:%s", err.Error())
		} else {
			logs.Error("当前连接没有ip，自动断开，断开连接成功")
		}
		return
	}
	if _, ok := clients.LoadOrStore(clientAddr, c); !ok {
		logs.Info("新客户端:%s", clientAddr)
	}
	defer c.Close()
	for {
		err = ws_common.WriteMessage(c, ws_common.MESSAGEPING, &protocol.Ping{
			PingVal: golibs.StandardTime(),
		})
		if err != nil {
			logs.Error("write:", err)
			clients.Delete(clientAddr)
			logs.Info("客户端:%s，退出", clientAddr)
			break
		}
		time.Sleep(time.Second * 10)
	}
}
