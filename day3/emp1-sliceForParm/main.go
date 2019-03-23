package main

import (
	"fmt"
	"math/rand"
	"time"
)

func initSlice(s []int)  {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(s) ;i++  {
		s[i]=rand.Intn(100)
	}
}

func bubbleSort(s []int)  {
	n:=len(s)
	for i:=0;i<n-1 ;i++  {
		for j:=0;j<n-1-i ;j++  {
			if s[j]>s[j+1] {
				s[j],s[j+1]=s[j+1],s[j]
			}
		}
	}
}

func main()  {

	n:= 10

	s:= make([]int,n)

	fmt.Println(s)

	initSlice(s)

	fmt.Println("初始化后的数组：",s)

	bubbleSort(s)

	fmt.Println("冒泡排序后的数组：",s)
}
