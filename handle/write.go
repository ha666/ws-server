package handle

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/ha666/logs"
	"github.com/ha666/ws-common/protocol"
)

func Write(c *websocket.Conn, dst proto.Message) {
	val, ok := dst.(*protocol.Write)
	if !ok {
		logs.Error("解析write消息出错")
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "write", val.WriteVal)
}
