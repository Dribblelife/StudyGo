package main

import (
	"net"
	"fmt"
	"strings"
)

func HandleConn(conn net.Conn) {
	//函数调用完自动关闭conn
	defer conn.Close()
	//获取客户端网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect successful")
	buf := make([]byte, 2048)
	//读取用户数据
	for   {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s:%s", addr, string(buf[:n]))
		//数据转换大写，发送给对应用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}
func main() {
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for   {
		conn,err:=listener.Accept()
		if err!=nil {
			fmt.Println(err)
			return
		}

		//处理用户发送的数据
		go HandleConn(conn)

	}
}
