package main

import (
	"fmt"
	"strconv"
)

func max(num1,num2 int)int  {
	if num1>num2 {
		return num1
	}else {
		return num2
	}
}

func main() {
	var a,b int
	fmt.Println("input a & b:")
	fmt.Scanf("%d%d",&a,&b)
	fmt.Println(max(a,b))

	var h string
	h="17"
	c,_:=strconv.ParseInt(h,8,32)
	fmt.Println(c)


}
