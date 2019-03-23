package main

import "fmt"

type student struct {
	name string
	id int
}


func main() {
	s:=make([]interface{},3)
	s[0]=1
	s[1]="abc"
	s[2]=student{"mac",666}

	for index,data:=range s{
		if value,ok:=data.(int);ok==true {
			fmt.Printf("s[%d]的类型是int，值为:%d\n",index,value)
		}else if value,ok:=data.(string);ok==true {
			fmt.Printf("s[%d]的类型是string，值为：%s\n",index,value)
		}else if value,ok:=data.(student);ok==true {
			fmt.Printf("s[%d]的类型是student，值为：name=%s,id=%d\n",index,value.name,value.id)
		}
	}

	for index,data:=range s{
		switch value:=data.(type) {
		case int:
			fmt.Printf("s[%d]的类型是int，值为:%d\n",index,value)
		case string:
			fmt.Printf("s[%d]的类型是string，值为：%s\n",index,value)
		case student:
			fmt.Printf("s[%d]的类型是student，值为：name=%s,id=%d\n",index,value.name,value.id)
		}
	}
}
