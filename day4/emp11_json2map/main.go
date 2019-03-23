package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonBuf:=`
	{
		"company": "Google",
		"subjects": [
		"go",
			"java",
			"c"
		],
		"isOK": true,
		"price": 66.66
	}`
	//创建一个map
	m:=make(map[string]interface{},4)
	err:=json.Unmarshal([]byte(jsonBuf),&m)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("m=%+v\n",m)
	//判断m中的数据类型，使用类型断言
	for key,value:=range m {
		switch data:=value.(type) {
		case string:
			fmt.Printf("map[%s]的类型为string，value = %s\n",key,data)
		case bool:
			fmt.Printf("map[%s]的类型为bool，value = %v\n",key,data)
		case float64:
			fmt.Printf("map[%s]的类型为float64，value = %f\n",key,data)
		case []string:
			fmt.Printf("map[%s]的类型为切片，value = %v\n",key,data)
		case interface{}:
			fmt.Printf("map[%s]的类型为interface{}，value = %v\n",key,data)
		}
	}



}
