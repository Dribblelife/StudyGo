package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	list:=os.Args	//获取命令行参数

	if len(list)!=3{
		fmt.Println("usage:xxx srcFileName dstFileName")
		return
	}
	srcFileName:=list[1]
	dstFileName:=list[2]
	if srcFileName==dstFileName {
		fmt.Println("源文件和目标文件名字不能相同")
		return
	}

	//以只读的方式打开源文件
	sF,err1:=os.Open(srcFileName)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	//新建目的文件
	dF,err2:=os.Create(dstFileName)
	if err2!=nil {
		fmt.Println(err2)
		return
	}

	//操作完成后，再关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心步骤，从源文件读取内容，往目的文件写，读多少写多少
	buf:=make([]byte,4*1024)//4k大小的临时缓冲区

	for   {
		n,err:=sF.Read(buf)
		if err!=nil{
			fmt.Println(err)
			if err == io.EOF {
				break
			}
		}
		dF.Write(buf[:n])
	}
}