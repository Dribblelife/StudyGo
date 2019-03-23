package main

import "fmt"

/*
	50个学生 存储其id
*/
func main()  {
	var id[50]int
	for i:=0;i<len(id);i++ {
		id[i]=i+1
		fmt.Printf("id[%d]=%d\n",i,id[i])
	}
}