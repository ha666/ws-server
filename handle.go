package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ha666/golibs"
	"github.com/ha666/logs"
	"github.com/ha666/ws-common"
	"github.com/ha666/ws-common/protocol"
	"github.com/ha666/ws-server/handle"
	"github.com/ha666/ws-server/service"
)

func process(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Error("upgrade:", err)
		return
	}
	clientAddr := service.ClientIpPort(r)
	if golibs.Length(clientAddr) <= 0 {
		if err = c.Close(); err != nil {
			logs.Error("当前连接没有ip，自动断开，断开连接失败:%s", err.Error())
		} else {
			logs.Error("当前连接没有ip，自动断开，断开连接成功")
		}
		return
	}
	if err = service.ClientHeartbeat(clientAddr, c); err != nil {
		logs.Error("心跳失败,%s:%s", clientAddr, err.Error())
		return
	}
	logs.Info("新客户端:%s", clientAddr)
	go read(clientAddr, r)
	go write(clientAddr)
}

func read(clientAddr string, r *http.Request) {
	for {
		val, ok := service.Clients.Load(clientAddr)
		if !ok || val == nil {
			logs.Error("连接%s已不存在", clientAddr)
			return
		}
		c, ok := val.(*websocket.Conn)
		if !ok || c == nil {
			logs.Error("连接%s解析失败", clientAddr)
			return
		}
		dst, messageType, err := ws_common.ReadMessage(c)
		if err != nil {
			logs.Error("read err,%s,err:%s", clientAddr, err.Error())
			if strings.Contains(err.Error(), "close 1006") {
				if err = service.CloseClient(clientAddr, c); err != nil {
					logs.Error("read,客户端%s退出失败:%s", clientAddr, err.Error())
				} else {
					logs.Error("read,客户端%s退出成功", clientAddr)
					return
				}
			}
			time.Sleep(3 * time.Second)
			continue
		}
		if bytes.Compare(messageType, ws_common.MESSAGEPING) == 0 {
			handle.Ping(c, dst, r)
		} else if bytes.Compare(messageType, ws_common.MESSAGEPONG) == 0 {
			handle.Pong(c, dst)
		} else if bytes.Compare(messageType, ws_common.MESSAGEREAD) == 0 {
			handle.Read(c, dst)
		} else if bytes.Compare(messageType, ws_common.MESSAGEWRITE) == 0 {
			handle.Write(c, dst)
		} else if bytes.Compare(messageType, ws_common.MESSAGESUBSCRIPTION) == 0 {
			handle.Subscription(c, dst)
		} else if bytes.Compare(messageType, ws_common.MESSAGEPUBLISH) == 0 {
			handle.Publish(c, dst)
		} else {
			logs.Error("无效的消息类型")
		}
	}
}

func write(clientAddr string) {
	for {
		time.Sleep(5 * time.Second)
		val, ok := service.Clients.Load(clientAddr)
		if !ok || val == nil {
			logs.Error("连接%s已不存在", clientAddr)
			return
		}
		c, ok := val.(*websocket.Conn)
		if !ok || c == nil {
			logs.Error("连接%s解析失败", clientAddr)
			return
		}
		if err := ws_common.WriteMessage(c, ws_common.MESSAGEPONG, &protocol.Pong{
			PongVal: fmt.Sprintf("时间：%s，客户端地址:%s", golibs.StandardTime(), clientAddr),
		}); err != nil {
			logs.Error("pong发送失败:%s", err.Error())
			if strings.Contains(err.Error(), "write: broken pipe") ||
				strings.Contains(err.Error(), "wsasend: An established connection was aborted by the software in your host machine.") {
				if err = service.CloseClient(clientAddr, c); err != nil {
					logs.Error("write,客户端%s退出失败:%s", clientAddr, err.Error())
				} else {
					logs.Error("write,客户端%s退出成功", clientAddr)
					return
				}
			}
		} else {

		}
	}
}
