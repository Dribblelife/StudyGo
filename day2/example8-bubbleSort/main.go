package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	/*a:= [10] int{1,2,6,4,2,1,5,2,8,2}
	n:=10*/
	rand.Seed(time.Now().UnixNano())
	var a=[10]int{}
	var n = len(a)
	fmt.Println("the array's element is:")
	for i:=0;i<n ;i++  {
		a[i]=rand.Intn(100)
		fmt.Printf("%d ",a[i])
	}
	fmt.Println()
	fmt.Println("the array is:\n",a)
	for i:=0;i<n-1 ;i++  {
		for j:=0;j<n-1-i ;j++  {
			if a[j]>a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	fmt.Println("the new array is:")
	for i:=0;i<n ;i++  {
		fmt.Printf("%d ", a[i])
	}
}
