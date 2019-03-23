package main

import (
	"net/http"
	"net"
	"time"
	"fmt"
)

var url=[]string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
	"http://www.woccc.sdf",
}

func main() {
	for _,value:=range url {
		c:=http.Client{
			Transport:&http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout:=time.Second*2
					return net.DialTimeout(network,addr,timeout)
				},
			},
		}
		response,err:=c.Head(value)
		if err!=nil {
			fmt.Println(err)
			continue 
		}
		fmt.Printf("head succ,status:%v\n",response.Status)
	}
	
}
