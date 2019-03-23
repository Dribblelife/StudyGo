package main

import (
	"strconv"
	"fmt"
)

func main() {
	slice:=make([]byte,0,100)
	slice=strconv.AppendBool(slice,true)
	slice=strconv.AppendInt(slice,666,10)
	slice=strconv.AppendQuote(slice,"abc")


	fmt.Println("slice=",string(slice))
	//其他类型转换成字符串
	var str string
	defer func() {
		if err:=recover();err!=nil {
			fmt.Println(err)
		}
	}()
	str=strconv.FormatBool(false)
	fmt.Println(str)
	str=strconv.FormatInt(12,10)
	fmt.Println(str)
	str=strconv.FormatFloat(22.22,'e',3,64)
	fmt.Println(str)
	str=strconv.Itoa(13)
	fmt.Println(str)
	//字符串转成其他类型
	var flag bool
	var err error
	flag,err=strconv.ParseBool("false")
	if err==nil{
		fmt.Println("flag=",flag)
	}else {
		fmt.Println("err=",err)
	}

	a,_:=strconv.Atoi("10")
	fmt.Println(a)
}
