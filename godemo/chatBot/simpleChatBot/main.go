package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	//准备从标准输入读取数据
	inputReader:=bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	//读取数据直到碰到\n为止
	input,err:=inputReader.ReadString('\n')
	if err!=nil{
		fmt.Printf("An error occurred:%s\n",err)
		//异常退出
		os.Exit(1)
	}else {
		//切片操作删除最后的\n
		userName:=input[:len(input)-1]
		fmt.Printf("Hello,%s!Can i help u?\n",userName)
	}
	for {
		input,err = inputReader.ReadString('\n')
		if err!=nil {
			fmt.Printf("An error occurred:%s\n",err)
			continue
		}
		input=input[:len(input)-1]
		input=strings.ToLower(input)		//输入内容转换为小写
		switch input {
		case "":
			continue
		case "nothing","bye":
			fmt.Println("ok,bye!")
			os.Exit(0)
		default:
			fmt.Println("sorry,i cannot catch u ")


		}
	}
}
