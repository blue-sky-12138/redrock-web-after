package database

import (
	"database/sql"
	"log"
	"os"
)

var(
	DataBase *sql.DB
)

//请在这里修改成自己的数据库信息
//连接数据库
func init(){
	temDatabase,err:=sql.Open("mysql","bluesky:135246Cjw@tcp(rm-bp14fk5x3q4byb6a2125010.mysql.rds.aliyuncs.com:3306)/redrock_web_data?charset=utf8&parseTime=True&loc=Local")
	DataBase =temDatabase
	CheckError("LinkMySQL error",err)
}

func CheckError(detail string,err error){
	logFile,_:=os.OpenFile("logFile" ,os.O_RDWR, os.ModePerm)
	defer logFile.Close()
	if err!=nil{
		detail="["+detail+"]"
		log.Printf(detail,err,"\n")

		debugLog := log.New(logFile,detail,log.Llongfile)
		debugLog.Println()
	}
}
