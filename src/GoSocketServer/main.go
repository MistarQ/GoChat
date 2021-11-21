package main

import (
	"chat/src/GoSocketServer/internal/config"
	"fmt"
	"time"
)

func main() {

	if err := config.Init("./src/GoSocketServer/config/config.dev.json") ; err !=nil {
		fmt.Println(err)
		return
	}

	// 启动Server
	go StartServer()
	fmt.Println("Go，Socket消息广播功能已启用")



	// 防止主线程退出
	for {
		time.Sleep(1 * time.Second)
	}
}

func StartServer() {
	server := NewServer(config.Conf.Server.Addr, config.Conf.Server.Port)
	server.Start()
}
