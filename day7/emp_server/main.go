package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()
	for  {
		buf:=make([]byte,1024)
		n,err:=conn.Read(buf)
		if err!=nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	fmt.Println("server starting...")
	listener,err:=net.Listen("tcp","0.0.0.0:50000")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	for   {
		conn,err:=listener.Accept()
		if err!=nil {
			fmt.Println(err)
			continue
		}
		go process(conn)
	}
}
