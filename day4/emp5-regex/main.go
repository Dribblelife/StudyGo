package main

import (
	"regexp"
	"fmt"
)

func main() {
	buf:="abc acc a8c 999 bac acc"
	//1)解析规则,他会解析正则表达式，如果成功则返回解析器
	reg1:=regexp.MustCompile(`a.c`)
	if reg1==nil {
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result:=reg1.FindAllStringSubmatch(buf,-1)
	fmt.Println(result)
}
