package main

import (
	"fmt"
	"os"
	"net"
	"io"
)

func SendFile(path string,conn net.Conn)  {
	//以只读的方式打开文件
	f,err:=os.Open(path)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	//发送文件内容
	buf:=make([]byte,4096)
	for   {
		n,err:=f.Read(buf)
		if err!=nil {
			if err==io.EOF {
				fmt.Println("文件发送完毕")
			}else{
				fmt.Println(err)
			}
			return
		}
		//给服务器发送内容
		conn.Write(buf[:n])
	}
}

func main() {
	//提示输入文件名

	fmt.Println("请输入文件名：")
	var path string
	fmt.Scan(&path)

	//获取文件名info.Name()

	info,err:=os.Stat(path)
	if err!=nil {
		fmt.Println(err)
		return
	}

	//主动连接server
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//给接收方发送文件名
	_,err	=conn.Write([]byte(info.Name()))
	if err!=nil {
		fmt.Println(err)
		return
	}

	//接收方接收对方的回复，如果回复“ok”，说明可以发送文件
	buf:=make([]byte,1024)
	n,err:=conn.Read(buf)
	if err!=nil {
		fmt.Println(err)
		return
	}
	if "ok"==string(buf[:n]) {
		SendFile(path,conn)
	}
}
