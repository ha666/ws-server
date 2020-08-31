package main

import (
	"github.com/ha666/logs"
	"sync"
	"time"
)

var clients sync.Map

func statisticsClientTotal() {
	for {
		time.Sleep(time.Second * 3)
		count:=0
		clients.Range(func(k,v interface{}) bool {
			count++
			return true
		})
		logs.Info("客户端总数:%d",count)
	}
}
