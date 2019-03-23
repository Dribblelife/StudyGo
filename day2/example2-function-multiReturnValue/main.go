package main

import (
	"fmt"
	)

func div(a,b int)(int,int)  {
	return a/b,a%b
}

func evaluate(a,b int ,op string)(int,error)  {
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		p,_:=div(a,b)
		return p,nil
	default:
		return 0,fmt.Errorf("unsupport operation "+ op)
	}
}

func main()  {
	c,d:=div(15,6)
	fmt.Println(c,d)
	//fmt.Println(evaluate(4,6,"/"))      			简洁写法
	if result,err :=evaluate(4,5,"x");err!=nil{		//常用写法
		fmt.Println("Error:",err)
	}else {
		fmt.Println(result)
	}
}
