package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	//主动链接服务器
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connect success")

	defer conn.Close()

	//从键盘输入内容，发送给服务器

	go func() {
		//创建一个缓冲切片
		str:=make([]byte,1024)
		for   {
			n,err:=os.Stdin.Read(str)	//从键盘读取内容，放在str
			if err != nil{
				fmt.Println(err)
				return
			}
			//把 输入的内容发送给服务器
			conn.Write(str[:n])
		}

	}()

	//接受服务器回复的数据，打印出来
	buf:=make([]byte,1024)	//创建一个缓冲切片
	for   {
		n,err:=conn.Read(	buf)
		if err!=nil {
			fmt.Println(err)
			return
		}
		fmt.Println("the data is :",string(buf[:n]))
	}
}
