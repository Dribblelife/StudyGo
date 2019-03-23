package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string `json:"company"`
	Subjects []string `json:"-"`
	IsOK bool	`json:",string"`
	Price float64 `json:",string"`		//结构体首字母必须大写？
}


func main() {
	jsonBuf:=`
		{
		"company": "Google",
		"subjects": [
				"go",
				"java",
				"c"
		],
		"isOK": "true",
		"price": "66.66"
}`
	var tmp IT
	err:=json.Unmarshal([]byte(jsonBuf),&tmp)
	if err!=nil {
		fmt.Println(err)
		return
	}else {
		fmt.Println(tmp)
		fmt.Printf("tmp is %+v\n",tmp)
	}
	type IT2 struct {
		Subjects []string `json:"subjects"`
	}
	var tmp2 IT2
	err =json.Unmarshal([]byte(jsonBuf),&tmp2)
	if err!=nil {
		fmt.Println(err)
		return
	}else {
		fmt.Printf("%+v\n",tmp2)
	}
}
