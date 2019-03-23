package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	)

func fileWrite(path string)  {
	f,err:=os.Create(path)
	if err!=nil {
		fmt.Println("err = ",err)
		return
	}
	defer f.Close()
	//通过defer，我们可以在代码中 延时关闭  优雅的关闭/清理代码中所使用的变量。这里是关闭文件

	var buf string
	for i:=0;i<10 ;i++  {
		//格式化字符串Sprintf
		buf=fmt.Sprintf("i=%d\n",i)
		fmt.Println("buf=",buf)
		n,err:=f.WriteString(buf)
		if err!=nil {
			fmt.Println("err = ",err)
		}
		fmt.Println("n = ",n)
	}

}

func FileRead(path string){
	f,err:=os.Open(path)
	if err!=nil {
		fmt.Println("err=",err)
		return
	}
	defer f.Close()

	buf:=make([]byte,1024*2)  	//每次读取2k大小

	n,err:=f.Read(buf)		//n代表从文件读取内容的长度
	if err!=nil && err==io.EOF {//文件错误或没有读到结尾
		fmt.Println("err=",err)
		return
	}
	fmt.Println("buf=",string(buf[:n]))
}

func ReadFilebyLine(path string){
	f,err:=os.Open(path)
	if err!=nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	r:=bufio.NewReader(f)
	for   {
		buf,err:=r.ReadString('\n')
		if err!=nil {
			if err== io.EOF{//文件已经结束
				break
		}
			fmt.Println(err)
		}
		fmt.Printf("buf=#%s#\n",string(buf))
	}

}

func main() {
	path:="./demo.txt"
	//fileWrite(path)
	//FileRead(path)
	ReadFilebyLine(path)
}
