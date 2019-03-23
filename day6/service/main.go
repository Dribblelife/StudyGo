package main

import (
	"net"
	"fmt"
)

func main() {

	//监听
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	//阻塞等待用户连接
	connect,err:=listener.Accept()
	if err!=nil {
		fmt.Println(err)
		return
	}
	//接受用户请求
	buf:=make([]byte,1024)//1024大小的缓冲区
	n,err:=connect.Read(buf)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("buf=",string(buf[:n]))
	defer connect.Close()
}
