package database

import (
	"database/sql"
	"log"
)

var(
	DataBase *sql.DB
)

//请在这里修改成自己的数据库信息
//连接数据库
func init(){
	temDatabase,err:=sql.Open("mysql","root:135246@tcp(localhost:3306)/redrock_homework6_web_users_information?charset=utf8&parseTime=True&loc=Local")
	DataBase =temDatabase
	CheckError("LinkMySQL error",err)
}

func CheckError(detail string,err error){
	if err!=nil{
		log.Printf("%s:", detail)
		println(err)
	}
}