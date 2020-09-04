package handle

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/ha666/logs"
	"github.com/ha666/ws-common/protocol"
)

func Publish(c *websocket.Conn, dst proto.Message) {
	val, ok := dst.(*protocol.Publish)
	if !ok {
		logs.Error("解析publish消息出错")
		return
	}
	logs.Info("\tmessageType:%s\tmessage: %s", "publish", val.PublishVal)
}
