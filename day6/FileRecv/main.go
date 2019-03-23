package main

import (
	"net"
	"fmt"

	"os"
	"io"
)

//接受文件
func FileRecv(filename string,conn net.Conn)  {
	//新建文件
	f,err:=os.Create(filename)
	if err!=nil {
		fmt.Println(err)
		return
	}
	//接受内容，读多少写多少
	buf:=make([]byte,4096)
	for   {
		n,err:=conn.Read(buf)
		if err!=nil {
			if err==io.EOF {
				fmt.Println("接受完毕")
			}else {
				fmt.Println(err)
			}
			return
		}
		f.Read(buf[:n])
	}
}

func main() {
	//监听
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	//阻塞用户连接
	conn,err:=listener.Accept()
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	buf:=make([]byte,1024)
	n,err:=conn.Read(buf)
	if err!=nil {
		fmt.Println(err)
		return
	}

	filename:=string(buf[:n])

	//回复ok
	conn.Write([]byte("ok"))

	//接受文件
	FileRecv(filename,conn)
}
