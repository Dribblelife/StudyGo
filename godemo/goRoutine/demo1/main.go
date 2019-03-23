package main

import (
	"fmt"
	"time"
)

func newTask()  {
	i:=0
	for {
		i++
		fmt.Printf("new goroutine is: i = %d\n",i)
		time.Sleep(1*time.Second)
	}
}

func main() {//创建一个goroutine 启动另一个任务
	go newTask()


	for  i:=0;i<10;i++ {
		fmt.Printf("main goroutine is: i = %d\n",i)
		time.Sleep(1*time.Second)
	}

	/*ips:=[]string{"1","2","3","4"}
	fmt.Println(ips[2:cap(ips)])*/
	/*result,err:=divide(2,2)
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}*/


}
/*func divide(a int,b int)(result int,err error)  {
	result=a/b
	return
}
*/