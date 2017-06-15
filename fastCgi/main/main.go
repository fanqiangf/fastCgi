package main

import (
	"fastCgi/common"
	"fastCgi/server"
	"fmt"
	"net"
)

func main() {
	service := common.GetRead("port")
	//绑定本地端口
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	common.CheckError(err, 17)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	common.CheckError(err, 19)
	fmt.Println("=====>")
	fmt.Println("启动socket监听")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go server.HandleClient(conn)

	}
}
