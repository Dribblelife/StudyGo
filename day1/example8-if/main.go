package main

import (
	"io/ioutil"
	"fmt"
)

/*func main()  {											常规if写法
	const filename  = "a.txt"
	content,err := ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s",content)
	}
}*/
func main()  {
	const filename="a.txt"
	if content,err :=ioutil.ReadFile(filename);err!=nil{	//简易if写法
		fmt.Println(err)
	}else {
		fmt.Printf("%s",content )
	}
}