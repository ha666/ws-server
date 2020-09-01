package main

import (
	"fmt"
	"github.com/ha666/golibs"
	"net/http"
	"os"
	"strings"
)

// 创建文件夹
func PathCreate(dir string) error {
	exist, err := PathExists(dir)
	if err != nil {
		return err
	}
	if exist {
		return nil
	} else {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
