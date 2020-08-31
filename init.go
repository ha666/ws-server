package main

import "github.com/ha666/logs"

func init() {
	initLog()
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
