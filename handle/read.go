package handle

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/ha666/logs"
	"github.com/ha666/ws-common/protocol"
)

func Read(c *websocket.Conn, dst proto.Message) {
	val, ok := dst.(*protocol.Read)
	if !ok {
		logs.Error("解析read消息出错")
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "read", val.ReadVal)
}
