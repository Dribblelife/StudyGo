package main

import (
	"regexp"
	"fmt"
)

func main()  {
	buf:="2.22 1.33 sad 555.234 99"
	//写一个正则表达式：+匹配前一个字符的一次或多次
	reg:=regexp.MustCompile(`\d*\.\d*`)
	if reg==nil{
		fmt.Println("regexp err")
	}

	result:=reg.FindAllStringSubmatch(buf,-1)
	fmt.Println(result)
}
