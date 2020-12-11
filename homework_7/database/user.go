package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Tem struct {
	TemString string
	TemInt    int
	TemSalt   int64
	TemBool bool
}

var (
	tem Tem
)

//寻找用户名
func FindUserName(ID int)string{
	prepare:=fmt.Sprintf("select name from users_information where telephone_number=%d",ID)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindUserID error",err)
	defer stmt.Close()
	if stmt.Next(){
		//存储用户名
		stmt.Scan(&tem.TemString)
		return tem.TemString
	}else{
		return ""
	}
}

//寻找用户手机号
func FindUserTelephone(name string)int{
	prepare:=fmt.Sprintf("select telephone_number from users_information where name='%s'",name)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindUserID error",err)
	defer stmt.Close()
	if stmt.Next(){
		//存储用户名
		stmt.Scan(&tem.TemInt)
		return tem.TemInt
	}else{
		return 0
	}
}

//寻找用户密码
func FindUserPassword(name string)(string,int64){
	prepare:=fmt.Sprintf("select password,MD5salt from users_information where name='%s'",name)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindUserPassword error",err)
	defer stmt.Close()
	if stmt.Next(){
		stmt.Scan(&tem.TemString,&tem.TemSalt)
	}
	return tem.TemString, tem.TemSalt
}

//寻找MD5盐
func FindSalt(ID int)int64{
	prepare:=fmt.Sprintf("select MD5salt from users_information where telephone_number=%d",ID)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindSalt error",err)
	defer stmt.Close()
	if stmt.Next(){
		stmt.Scan(&tem.TemSalt)
	}
	return tem.TemSalt
}

//执行更改操作
func Update(ID int,target string,detail string){
	prepare:=fmt.Sprintf("update users_information set %s='%s' where telephone_number=%d",target,detail,ID)
	stmt,err:= DataBase.Prepare(prepare)
	defer stmt.Close()
	CheckError("Update error",err)
	stmt.Exec()
}

//添加注册数据
func InsertUser(name string,password string,telephoneNumber int,MD5salt int64){
	prepare:=fmt.Sprintf("insert users_information (name,password,telephone_number,MD5salt)value('%s','%s',%d,%d)",name,password,telephoneNumber,MD5salt)
	stmt,err:= DataBase.Prepare(prepare)
	defer stmt.Close()
	CheckError("InsertUser error",err)
	stmt.Exec()
}

func FindUserId(name string) int {
	prepare:=fmt.Sprintf("select id from users_information where name='%s'",name)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindUserId error",err)
	defer stmt.Close()
	if stmt.Next(){
		stmt.Scan(&tem.TemInt)
	}
	return tem.TemInt
}