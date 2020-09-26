package main

import (
	"github.com/google/gops/agent"
	"github.com/ha666/logs"
	"github.com/ha666/ws-server/service"
	"log"
)

func init() {
	initLog()
	initGops()
	service.InitRedis()
}

func initLog() {
	if err := PathCreate("./log"); err != nil {
		logs.Emergency("创建目录出错:%s", err.Error())
	}
	_ = logs.SetLogger(logs.AdapterConsole, `{"level":0}`)
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"./log/log.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":100}`)
	//_ = logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}

func initGops() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
}
