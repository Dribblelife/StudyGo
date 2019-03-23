package main

import (
	"fmt"
	)

func main()  {
	var a[5] int			//定义数组元素个数必须是常量
	var b[10] int
	fmt.Printf("len(a)=%d,len(b)=%d\n",len(a),len(b))

	for i:=0;i<len(a) ;i++  {
		a[i]=i+1
	}
	for i,data:=range a{				//range 遍历很有用
		fmt.Printf("a[%d]=%d\n",i,data)
	}
}
