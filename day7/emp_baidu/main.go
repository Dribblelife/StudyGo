package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	conn,err:=net.Dial("tcp","www.baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	msg:="GET / HTTP/1.1\r\n"
	msg +="Host:www.baidu.com\r\n"
	msg +="Connection:keep-alive\r\n"
	msg +="\r\n\r\n"
	n,err:=io.WriteString(conn,msg)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send to baidu.com bytes:",n)
	buf:=make([]byte,1024)
	for  {
		count,err:=conn.Read(buf)
		fmt.Println("count:",count,"err:",err)
		if err!=nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:count]))
	}
	}
