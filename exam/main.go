package main

import (
	"fmt"
)

func sameNums(a[] int,b[] int)  {
	var max,min int
	//判断谁的边界更大或更小
	if (a[len(a)-1]>b[len(b)-1]) {
		max=a[len(a)-1]
	}else {
		max=b[len(b)-1]
	}
	if (a[0]<b[0]) {
		min=a[0]
	}else {
		min=b[0]
	}

	c:=make([]int,max-min+1)
	//创建切片c，作为输出结果
	var leng int
	//设a，b长度最大值
	if len(a)<=len(b) {
		leng=len(b)
	}else {
		leng=len(a)
	}
	for i:=0;i<leng;i++  {
		var c1  = c[a[i]]-min
		var c2  = c[b[i]]-min
		if i<len(a) {
			if c1==0 {
				c1=a[i]
			}else {
				fmt.Println(c1)
			}
		}
		if i<len(b) {
			if c2==0 {
				c2=b[i]
			}else {
				fmt.Println(c2)
			}
		}
		fmt.Println(c)
	}

}

func main()  {
	var a,b []int
	fmt.Println("please input your array a:")
	fmt.Scanln(a)
	fmt.Println("please input your array b:")
	fmt.Scanln(b)
	sameNums(a,b)
	/*inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your array:(like{1,2,3})")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return*/

}