package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	name string
	password string
	studentNumber int
}

func main(){
	//请在这里修改成自己的数据库信息
	dataBase,err:=sql.Open("mysql","root:135246@tcp(localhost:3306)/redrockhomework?charset=utf8&parseTime=True&loc=Local")
	CheckError(err)

	var user User

	//登录
	fmt.Println("欢迎您使用自助服务系统，请输入你的账号和密码")
	for{
		fmt.Printf("您的账号(学生号)：")
		fmt.Scanf("%d",&user.studentNumber)
		if checkUser(dataBase,user.studentNumber){
			break
		}else{
			fmt.Println("该账号不存在！")
		}
	}
	fmt.Println("************************************")
	for{
		fmt.Printf("请输入您的密码：")
		fmt.Scanf("%s",&user.password)
		if user.checkPassword(dataBase){
			fmt.Println("************************************")
			fmt.Printf("欢迎回来%s\n",user.name)
			break
		}else{
			fmt.Println("密码错误！")
		}
	}

	if user.name=="Administrator"{
		for{
			fmt.Println("请选择你的操作：1.增添新用户信息  2.删除用户信息")
			fmt.Println("             3.修改用户密码    4.查询用户信息")
			fmt.Println("             5.退出系统")
			choice:=checkChoice(1,4)
			switch choice {
			case 1:
				insert(dataBase)
			case 2:
				delete(dataBase)
			case 3:
				user.update(dataBase)
			case 4:
				user.find(dataBase)
			}
			fmt.Println("************************************")
			if choice==5{
				fmt.Println("欢迎再次使用")
				break
			}
		}
	}else{
		for{
			fmt.Println("请选择你的操作：1.修改密码  2.退出系统")
			choice:=checkChoice(1,2)
			if choice==1{
				user.update(dataBase)
			}else{
				fmt.Println("************************************")
				fmt.Println("欢迎再次使用")
				break
			}
			fmt.Println("************************************")
		}
	}
}

func insert(database *sql.DB){
	fmt.Println("************************************")
	fmt.Println("请输入增添信息")
	var temInformation User
	fmt.Printf("用户名:")
	fmt.Scanf("%s",&temInformation.name)
	fmt.Printf("密码:")
	fmt.Scanf("%s",&temInformation.password)
	fmt.Printf("学生号:")
	fmt.Scanf("%d",&temInformation.studentNumber)
	prepare:=fmt.Sprintf("insert into homeworktable (name,password,studentnumber)values ('%s','%s',%d)",temInformation.name,temInformation.password,temInformation.studentNumber)
	stmt,err:=database.Prepare(prepare)
	CheckError(err)
	stmt.Exec()
	fmt.Println("执行完毕")
}

func (user User)update(database *sql.DB){
	var (
		prepare string
		temInformation User
	)
	if user.name=="Administrator"{
		fmt.Println("************************************")
		fmt.Printf("请输入修改目标学生号:")
		fmt.Scanf("%d",&temInformation.studentNumber)
		fmt.Printf("要更改的密码:")
		fmt.Scanf("%s",&temInformation.password)
		prepare=fmt.Sprintf("update homeworktable set password='%s' where studentnumber=%d",temInformation.password,temInformation.studentNumber)
	}else{
		fmt.Println("************************************")
		fmt.Printf("要更改的密码:")
		fmt.Scanf("%s",&temInformation.password)
		prepare=fmt.Sprintf("update homeworktable set password='%s' where studentnumber=%d",temInformation.password,user.studentNumber)
	}
	stmt,err:=database.Prepare(prepare)
	CheckError(err)
	stmt.Exec()
	fmt.Println("执行完毕")
}

func delete(database *sql.DB){
	fmt.Println("************************************")
	fmt.Printf("请输入要删除的目标学生号:")
	var studentNumber int
	fmt.Scanf("%d",&studentNumber)
	stmt,err:=database.Prepare("delete from homeworktable where studentnumber=?")
	CheckError(err)
	stmt.Exec(studentNumber)
	fmt.Println("删除完毕")
}

func (user User)find(database *sql.DB){
	var temNumber int
	if user.name=="Administrator"{
		fmt.Println("************************************")
		fmt.Printf("请输入查找的目标学生号:")
		fmt.Scanf("%d",&temNumber)
	}else{
		temNumber=user.studentNumber
	}
	prepare:=fmt.Sprintf("select * from homeworktable where studentnumber=%d",temNumber)
	stmt,err:=database.Query(prepare)
	CheckError(err)
	defer stmt.Close()
	var (
		temInformation User
		temGarbage int
	)
	for stmt.Next(){
		stmt.Scan(&temGarbage,&temInformation.name,&temInformation.password,&temInformation.studentNumber)
	}
	fmt.Printf("查询信息如下:\n")
	fmt.Printf("用户名:%s\n",temInformation.name)
	fmt.Printf("密码:%s\n",temInformation.password)
}

//检查用户是否存在
func checkUser(database *sql.DB,studentNumber int)bool{
	stmt,err:=database.Query("select password from homeworktable where studentnumber=?",studentNumber)
	CheckError(err)
	defer stmt.Close()
	if stmt.Next(){
		return true
	}
	return false
}

//检查密码是否正确
func (user *User)checkPassword(database *sql.DB)bool{
	stmt,err:=database.Query("select name,password from homeworktable where studentnumber=?",user.studentNumber)
	CheckError(err)
	defer stmt.Close()
	var tem string
	for stmt.Next(){
		stmt.Scan(&user.name,&tem)
	}
	if user.password==tem{
		return true
	}
	return false
}

func CheckError(err error){
	if err!=nil{
		log.Println(err)
	}
}

//录入用户选择，判断是否在整数a与b之间，并自动排除错误
func checkChoice(left,right int) int {
	var (
		xuanZhe =0
		ret int
	)
	for{
		ret,_=fmt.Scanf("%d",&xuanZhe)
		//检查是否成功录入数据
		if ret==0{
			fmt.Printf("您的选择不在可选项中\n")
			fmt.Scanf("%s")
		}
		if xuanZhe >= left && xuanZhe <= right{
			break
		}else{
			fmt.Printf("您的选择不在可选项中\n")
		}
	}
	return xuanZhe
}