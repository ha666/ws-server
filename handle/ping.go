package handle

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/ha666/golibs"
	"github.com/ha666/logs"
	"github.com/ha666/ws-common/protocol"
	"github.com/ha666/ws-server/service"
)

func Ping(c *websocket.Conn, dst proto.Message, r *http.Request) {
	val, ok := dst.(*protocol.Ping)
	if !ok {
		logs.Error("解析ping消息出错")
		return
	}
	clientAddr := service.ClientIpPort(r)
	if golibs.Length(clientAddr) <= 0 {
		if err := c.Close(); err != nil {
			logs.Error("当前连接没有ip，自动断开，断开连接失败:%s", err.Error())
		} else {
			logs.Error("当前连接没有ip，自动断开，断开连接成功")
		}
		return
	}
	if err := service.ClientHeartbeat(clientAddr, c); err != nil {
		logs.Error("心跳失败,%s:%s", clientAddr, err.Error())
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "ping", val.PingVal)
}
