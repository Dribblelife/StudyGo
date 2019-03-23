package main

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

var DB *sqlx.DB

func init()  {
	database,err:=sqlx.Open("mysql","root:Mac_2019@tcp(localhost:3306)/test")	//连接数据库
	if err!=nil {
		fmt.Println("open sql failed",err)
		return
	}
	DB=database
}

func main() {
	r,err:=DB.Exec("insert into person(username,sex,email)values (?,?,?)","stu001","man","stu1@qq.com")
	if err!=nil {
		fmt.Println("execute failed",err)
		return
	}
	id,err:=r.LastInsertId()
	if err!=nil {
		fmt.Println("execute failed",err)
		return
	}
	fmt.Println("insert success",id)
}
