package main

import "fmt"

func add(a int,arg... int)int  {
	var sum int = a
	for i:=0;i<len(arg);i++{
		sum+=arg[i]
	}
	return sum
}

/*
		下面这个是字符拼接
*/

func concat1(a string,arg... string)string  {		//自己想的方法
	var result string = a
	for i:=0;i<len(arg);i++{
		result+=arg[i]
	}
	return result
}

func concat2(a string ,arg... string)(result string){		//老师给的方法
	result=a
	for i:=0;i<len(arg) ;i++  {
		result+=arg[i]
	}
	return
}

func main()  {
	fmt.Println(add(15,3,2,11,3))
	fmt.Println(concat1("dad","f","fuck","cart"))

	resul:=concat2("mac","is","the","best")
	fmt.Println(resul)
}
