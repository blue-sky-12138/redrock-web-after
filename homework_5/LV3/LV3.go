package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
type drawContext struct{			//抽奖页面内容结构体
	Name string
	Brief string
	Path string
}
var nameChan=make(chan int,1)	//输送用户名，达到在两个函数中传递数据
var registerWrongChan=make(chan string,1)	//输送注册错误信息

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
	router.LoadHTMLGlob("templates/html/*")
	router.Static("/css", "./templates/css")
	router.Static("/image", "./templates/image")

	//显示登录界面
	router.GET("/Login", func(context *gin.Context) {
		context.HTML(http.StatusOK,"login.html",nil)
	})
	//接受传入登录数据
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
			context.Redirect(http.StatusMovedPermanently,"/LoginFalse1")
		}else if err==2{
			context.Redirect(http.StatusMovedPermanently,"/LoginFalse2")
		}
	})
	//登录检测结果
	router.GET("/LoginFalse1", func(c *gin.Context) {
		c.String(http.StatusOK,"该用户名不存在")
	})
	router.GET("/LoginFalse2", func(c *gin.Context) {
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
		switch {
		case err=="":
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
		err:=<-registerWrongChan
		c.String(http.StatusOK,"注册失败，"+err)
	})

	//上传页面
	router.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK,"upload.html",nil)
	})

	//抽奖页面
	router.GET("/draw", func(context *gin.Context) {
		context.HTML(http.StatusOK,"draw.html",nil)
	})
	router.GET("/drawBack", func(context *gin.Context) {
		var nameTem,briefTem,pathTem string
		rand.Seed(time.Now().Unix())		//随机种子
		switch rand.Intn(2){
		case 0:{
			nameTem="星星学姐"
			briefTem="超人气学姐，web红岩女团成员之一"
			pathTem="./templates/image/1.jpg"
		}
		case 1:{
			nameTem="鑫鑫学姐"
			briefTem="高冷气质学姐，web红岩女团成员之一"
			pathTem="./templates/image/2.jpg"
		}
		case 2:{
			nameTem="峰峰学姐"
			briefTem="温柔派学姐，web红岩女团成员之一"
			pathTem="./templates/image/3.jpg"
		}
		}
		drawBack:=drawContext{
			Name: nameTem,
			Brief:briefTem ,
			Path:pathTem,
		}
		context.HTML(http.StatusOK,"drawBack.html",drawBack)
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
func RegisterIn(nameKeyMap map[int]string,register Register,f *os.File)string{
	_,pd :=nameKeyMap[register.name]
	if register.name<=10000000000 || register.name>=100000000000{
		//简易判断手机号位数
		return "手机号错误"
	}else if pd==true{
		return "用户名已存在"
	}

	//密码输入
	if strings.Contains(register.password,"/"){
		//简易判断密码非法符号
		return "密码含有非法字符"
	}else if len(register.password)<=6{
		return "密码过短"
	}
	if register.passwordAgain!=register.password{
		return "两次密码不一致"
	}
	newInformation:=fmt.Sprintf("%d %s\n",register.name,register.password)
	n,_:=f.Seek(0,2)
	f.WriteAt([]byte(newInformation),n)
	return ""
}