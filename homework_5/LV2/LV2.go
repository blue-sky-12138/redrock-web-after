package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Login struct{
	name int
	password string
}
type Register struct{
	name int
	password string
	passwordAgain string
}
var nameChan=make(chan int,1)	//输送用户名，达到在两个函数中传递数据
var registerWrongChan=make(chan int,1)	//输送注册错误信息

//http://localhost:8080/index
func main() {
	//加载用户数据
	contact := make([]byte, 0)
	nameKeyMap :=make(map[int]string)
	f,_:=os.OpenFile("users.data" ,os.O_RDWR, os.ModePerm)
	defer f.Close()
	contact, err := ioutil.ReadFile("users.data")
	if err != nil {
		panic(err)
	}
	//读取文件中的用户名和密码并存入map
	nameKey:=bytes.Split(contact,[]byte("\n"))
	for _,x:=range nameKey{
		if len(x)==0{
			break
		}
		if x[len(x)-1] == 13 {
			x=x[:len(x)-1]					//去除密码中的\n
		}
		nameKeySplit:=bytes.Split(x,[]byte(" "))	//分开用户名和密码
		name,_:=strconv.Atoi(string(nameKeySplit[0]))	//存入账号
		nameKeyMap[name]=string(nameKeySplit[1])	//存入map
	}


	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/css", "./templates/css")

	//显示登录界面
	router.GET("/Login", func(context *gin.Context) {
		context.HTML(http.StatusOK,"login.html",nil)
	})
	//接受传入数据
	router.POST("/Login/Create", func(context *gin.Context) {
		var login Login
		name, _ :=context.GetPostForm("name")
		login.name,_= strconv.Atoi(name)
		login.password, _ =context.GetPostForm("password")

		//登录检测并重定向
		err,ifSuccess:=logIn(nameKeyMap,login.name,login.password)
		if ifSuccess{
			nameChan <-login.name
			name:= strconv.Itoa(<-nameChan)		//int转字符串
			cookie := &http.Cookie{
				Name:     "users",
				Value:    name,
				MaxAge: 1000000,
				Path:     "/",
				Domain: "localhost",
				Secure: false,
				HttpOnly: true,
			}
			http.SetCookie(context.Writer,cookie)
			context.Redirect(http.StatusMovedPermanently,"/index")
		}else if err==1{
			context.Request.URL.Path= "/LoginFalse1"
			router.HandleContext(context)
		}else if err==2{
			context.Request.URL.Path= "/LoginFalse2"
			router.HandleContext(context)
		}
	})
	//登录检测结果
	router.POST("/LoginFalse1", func(c *gin.Context) {
		c.String(http.StatusOK,"该用户名不存在")
	})
	router.POST("/LoginFalse2", func(c *gin.Context) {
		c.String(http.StatusOK,"密码错误")
	})

	router.GET("/index", func(context *gin.Context) {
		cookie, err :=context.Cookie("users")
		if err==nil{
			context.HTML(http.StatusOK,"index.html","欢迎回来"+cookie)
		}else{
			context.HTML(http.StatusOK,"index.html","你好游客")
		}
	})

	//显示注册界面
	router.GET("/Register", func(context *gin.Context) {
		context.HTML(http.StatusOK,"register.html",nil)
	})
	router.POST("/Register/Create", func(context *gin.Context) {
		var register Register
		name, _ :=context.GetPostForm("name")
		register.name,_= strconv.Atoi(name)
		register.password, _ =context.GetPostForm("password")
		register.passwordAgain, _ =context.GetPostForm("passwordAgain")

		//注册检测并重定向
		err:=RegisterIn(nameKeyMap,register,f)
		registerWrongChan<-err
		switch err{
		case 0:
			context.Redirect(http.StatusMovedPermanently,"/RegisterSuccess")
		default:
			context.Redirect(http.StatusMovedPermanently,"/RegisterFalse")
		}
	})
	//注册检测结果
	router.GET("/RegisterSuccess", func(c *gin.Context) {
		c.String(http.StatusOK,"注册成功")
	})
	router.GET("/RegisterFalse", func(c *gin.Context) {
		err:=strconv.Itoa(<-registerWrongChan)
		c.String(http.StatusOK,"注册失败，问题代号："+err+"\n代号全解：\n1.手机号错误\n2.用户名已存在\n3.密码含有非法字符\n4.密码过短\n5.两次密码不一致")
	})


	router.Run()
}


//登录，返回错误类型和是否登陆成功
func logIn(nameKeyMap map[int]string ,name int,key string)(int,bool){
	//搜索用户名
	for {
		_,pd :=nameKeyMap[name]
		if pd{
			break
		}else{
			return 1,false
		}
	}

	//验证密码
	for{
		yz,_ :=nameKeyMap[name]
		if yz==key{
			return 0,true					//用于退出系统的指示变量
		}else{
			return 2,false
		}
	}
}

//注册
func RegisterIn(nameKeyMap map[int]string,register Register,f *os.File)int{
	_,pd :=nameKeyMap[register.name]
	if register.name<=10000000000 || register.name>=100000000000{
		//简易判断手机号位数
		return 1
	}else if pd==true{
		return 2
	}

	//密码输入
	if strings.Contains(register.password,"/"){
		//简易判断密码非法符号
		return 3
	}else if len(register.password)<=6{
		return 4
	}
	if register.passwordAgain!=register.password{
		return 5
	}
	newInformation:=fmt.Sprintf("%d %s\n",register.name,register.password)
	n,_:=f.Seek(0,2)
	f.WriteAt([]byte(newInformation),n)
	return 0
}