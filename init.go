package main

import (
	"time"

	"github.com/ha666/golibs"
	"github.com/ha666/logs"
	"gopkg.in/redis.v5"
)

var (
	redisClient *redis.Client
)

func init() {
	initLog()
	initRedis()
}

func initLog() {
	if err := PathCreate("./log"); err != nil {
		logs.Emergency("创建目录出错:%s", err.Error())
	}
	_ = logs.SetLogger(logs.AdapterConsole, `{"level":0}`)
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"./log/log.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":100}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:       "127.0.0.1:6379",
		Password:   "",
		PoolSize:   20,
		MaxRetries: 5,
	})
	err := redisClient.Ping().Err()
	if err != nil {
		logs.Emergency("initRedis错误:%s", err.Error())
	}
	guid := golibs.GetGuid()
	setResult := redisClient.Set("initRedis-guid", guid, 10*time.Second)
	if setResult.Err() != nil {
		logs.Emergency("initRedis初始化,写入guid出错:%s", setResult.Err().Error())
	}
	getResult := redisClient.Get("initRedis-guid")
	if getResult.Err() != nil {
		logs.Emergency("initRedis初始化,查询guid出错:%s", getResult.Err().Error())
	}
	if getResult.Val() != guid {
		logs.Emergency("initRedis初始化,对比guid不通过")
	}
	logs.Info("redis初始化成功")
}
