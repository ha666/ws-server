package main

import (
	"os"
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
