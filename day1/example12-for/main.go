package main

import "fmt"

func main()  {
	print(5)
	i:=5
	for {
		if i<=0{
			break
		}
		fmt.Println("",i)
		i--
	}
	for i:=0;i<10;i++{
		if i%2==0{
			continue
		}
		fmt.Println("",i)
		i++
	}
}

func print(n int)  {
	for i:=0;i<n;i++  {
		for j:=0;j<i+1;j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}