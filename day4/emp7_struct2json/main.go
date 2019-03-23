package main

import (
			"fmt"
	"encoding/json"
)

type IT struct {
	Company string `json:"company"` //变小写
	Subjects []string `json:"-"` //此字段不会输出到屏幕
	IsOk bool `json:",string"`	//变成字符串输出
	Money float64 `json:",string"`
}


func main() {
	var s=IT{"huawei",[]string{"go","web","bigdata"},false,15000.000}
	//buf,err:=json.Marshal(s)
	buf,err:=json.MarshalIndent(s,"","	")
	if err!=nil {
		fmt.Println("err is ",err)
		return
	}else {
		fmt.Println(string(buf))
	}
}
