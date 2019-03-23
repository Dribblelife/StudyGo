package main

import (
	"fmt"
	"errors"
)

func main() {
	err1:=fmt.Errorf("%s","this is normal error")
	fmt.Println(err1)

	err2:=errors.New("this is normal err2")
	fmt.Println(err2)

	result,err:=divide(10,0)
	if err!=nil{
		fmt.Println("err is \n",err)
	}else {
		fmt.Println("the result is \n",result)
	}
}

func divide(a,b float64) (result float64,err error) {
	err=nil
	if b==0{
		err=errors.New("分母不能为0")
	}else {
		result=a/b
	}
	return
}