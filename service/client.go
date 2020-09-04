package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ha666/golibs"
	"github.com/ha666/logs"
	"gopkg.in/redis.v5"
)

var (
	Clients     sync.Map
	RedisClient *redis.Client
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:       "127.0.0.1:6379",
		Password:   "",
		PoolSize:   20,
		MaxRetries: 5,
	})
	err := RedisClient.Ping().Err()
	if err != nil {
		logs.Emergency("initRedis错误:%s", err.Error())
	}
	guid := golibs.GetGuid()
	setResult := RedisClient.Set("initRedis-guid", guid, 10*time.Second)
	if setResult.Err() != nil {
		logs.Emergency("initRedis初始化,写入guid出错:%s", setResult.Err().Error())
	}
	getResult := RedisClient.Get("initRedis-guid")
	if getResult.Err() != nil {
		logs.Emergency("initRedis初始化,查询guid出错:%s", getResult.Err().Error())
	}
	if getResult.Val() != guid {
		logs.Emergency("initRedis初始化,对比guid不通过")
	}
	logs.Info("redis初始化成功")
}

func StatisticsClientTotal() {
	for {
		time.Sleep(time.Second * 3)
		count := 0
		Clients.Range(func(k, v interface{}) bool {
			count++
			return true
		})
		logs.Info("客户端总数:%d", count)
	}
}

//客户端心跳
func ClientHeartbeat(clientAddr string, c *websocket.Conn) error {
	Clients.LoadOrStore(clientAddr, c)
	result := RedisClient.ZAdd("clients", redis.Z{
		Score:  float64(golibs.UnixMilliSecond()),
		Member: clientAddr,
	})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

//客户端关闭
func ClientClose(clientAddr string) error {
	Clients.Delete(clientAddr)
	result := RedisClient.ZRem("clients", clientAddr)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func GetClient(clientAddr string) (*websocket.Conn, error) {
	val, ok := Clients.Load(clientAddr)
	if !ok || val == nil {
		return nil, errors.New("没有找到连接")
	}
	c, ok := val.(*websocket.Conn)
	if !ok || val == nil {
		return nil, errors.New("没有找到连接")
	}
	return c, nil
}

func ClientIpPort(r *http.Request) string {
	ip := ""
	port := ""
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if golibs.Length(xForwardedFor) <= 0 {
		return r.RemoteAddr
	}
	ip = strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if golibs.Length(ip) <= 0 {
		return ""
	}
	port = strings.TrimSpace(r.Header.Get("Remote-Port"))
	if golibs.Length(port) <= 0 {
		return ""
	}
	return fmt.Sprintf("%s:%s", ip, port)
}
