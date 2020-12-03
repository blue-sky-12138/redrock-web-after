package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID int
	name string
	password string
}
type temGarbage struct {	//用于储存搜索数据库时产生的无用数据
	intGarbage int
	stringGarbage string
}
type Change struct {		//用于储存用户修改的信息
	ChangeName string
	ChangeNameAgain string
	ChangePassword string
	ChangePasswordAgain string
	Signature string
}
var (
	Garbage temGarbage
    login User
	PastContent struct{Content string}
	cookie *http.Cookie
)

func main(){
	//请在这里修改成自己的数据库信息
	//连接数据库
	dataBase,err:=sql.Open("mysql","root:135246@tcp(localhost:3306)/redrock_homework6_web_users_information?charset=utf8&parseTime=True&loc=Local")
	CheckError(err)

	//http://localhost:8080/index
	//启动服务器
	router := gin.Default()
	router.LoadHTMLGlob("templates/html/*")
	router.Static("/css", "./templates/css")

	//登录界面
	router.GET("/Login", func(context *gin.Context) {
		context.HTML(http.StatusOK,"login.html",nil)
	})
	//接受传入登录数据
	router.POST("/Login/Create", func(context *gin.Context) {
		temName, _ :=context.GetPostForm("name")
		login.ID,_= strconv.Atoi(temName)
		login.password,_=context.GetPostForm("password")

		//登录检测并重定向
		if login.checkUser(dataBase) && login.checkPassword(dataBase){
			cookie := &http.Cookie{
				Name:     "users",
				Value:    login.name,
				MaxAge: 100000,
				Path:     "/",
				Domain: "localhost",
				Secure: false,
				HttpOnly: true,
			}
			http.SetCookie(context.Writer,cookie)
			context.Redirect(http.StatusMovedPermanently,"/index")
		}else{
			context.Redirect(http.StatusMovedPermanently,"/LoginFalse")
		}
	})
	//登录检测结果
	router.GET("/LoginFalse", func(c *gin.Context) {
		c.String(http.StatusOK,"请检查你的用户名和密码是否正确")
	})

	//主页
	router.GET("/index", func(context *gin.Context) {
		cookie, err :=context.Cookie("users")
		var information struct{
			Title string
			VisterOK bool
			LoginOK bool
		}
		if err==nil{
			if login.ID==0{
				login.name=cookie
				login.findUserInformation(dataBase)
			}
			information.Title="欢迎回来"+login.name
			information.LoginOK=true
			information.VisterOK=false
			context.HTML(http.StatusOK,"index.html",information)
		}else{
			information.Title="你好游客"
			information.LoginOK=false
			information.VisterOK=true
			context.HTML(http.StatusOK,"index.html",information)
		}
	})

	//更改名称
	router.GET("/ChangeName", func(context *gin.Context) {
		context.HTML(http.StatusOK,"ChangeName.html",nil)
	})
	router.POST("/ChangeName/Create", func(context *gin.Context) {
		var tem Change
		tem.ChangeName,_=context.GetPostForm("ChangeName")
		tem.ChangeNameAgain,_=context.GetPostForm("ChangeNameAgain")
		if tem.ChangeName!=tem.ChangeNameAgain{
			PastContent.Content="两次输入的名称不一致！"
			context.Redirect(http.StatusMovedPermanently,"/Past")
		}else{
			cookie.Value=tem.ChangeName
			login.update(dataBase,"name",tem.ChangeName)
			PastContent.Content="更改完成"
			context.Redirect(http.StatusMovedPermanently,"/Past")
		}
	})

	//更改密码
	router.GET("/ChangePassword", func(context *gin.Context) {
		context.HTML(http.StatusOK,"ChangePassword.html",nil)
	})
	router.POST("/ChangePassword/Create", func(context *gin.Context) {
		var tem Change
		tem.ChangePassword,_=context.GetPostForm("ChangePassword")
		tem.ChangePasswordAgain,_=context.GetPostForm("ChangePasswordAgain")
		if tem.ChangePassword!=tem.ChangePasswordAgain{
			PastContent.Content="两次输入的密码不一致！"
			context.Redirect(http.StatusMovedPermanently,"/Past")
		}else{
			login.update(dataBase,"password",tem.ChangePassword)
			PastContent.Content="更改完成"
			context.Redirect(http.StatusMovedPermanently,"/Past")
		}
	})

	//更改个性签名
	router.GET("/ChangeSignature", func(context *gin.Context) {
		context.HTML(http.StatusOK,"ChangeSignature.html",nil)
	})
	router.POST("/ChangeSignature/Create", func(context *gin.Context) {
		var tem Change
		tem.Signature,_=context.GetPostForm("ChangeSignature")
		login.update(dataBase,"signature",tem.Signature)
		PastContent.Content="更改完成"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	})

	//过渡页面
	router.GET("/Past", func(context *gin.Context) {
		context.HTML(http.StatusOK,"Past.html",PastContent)
	})

	router.Run(":8080")
}

//执行更改操作
func (user *User)update(database *sql.DB,target string,detail string){
	prepare:=fmt.Sprintf("update users_information set %s='%s' where telephone_number=%d",target,detail,user.ID)
	stmt,err:=database.Prepare(prepare)
	defer stmt.Close()
	CheckError(err)
	stmt.Exec()
}

//检查用户是否存在
func (user *User)checkUser(database *sql.DB)bool{
	prepare:=fmt.Sprintf("select name from users_information where telephone_number=%d",user.ID)
	stmt,err:=database.Query(prepare)
	CheckError(err)
	defer stmt.Close()
	if stmt.Next(){
		//存储用户名
		stmt.Scan(&user.name)
		return true
	}
	return false
}

//检查密码是否正确
func (user *User)checkPassword(database *sql.DB)bool{
	prepare:=fmt.Sprintf("select password from users_information where name='%s'",user.name)
	stmt,err:=database.Query(prepare)
	CheckError(err)
	defer stmt.Close()
	var tem string
	for stmt.Next(){
		stmt.Scan(&tem)
	}
	if user.password==tem{
		return true
	}
	return false
}

//解决有cookie时用户的数据(这里不考虑安全性问题，我还不会使用cookie存储更多信息，我只会直接使用其value值)
func (user *User)findUserInformation(database *sql.DB){
	prepare:=fmt.Sprintf("select telephone_number from users_information where name='%s'",user.name)
	stmt,err:=database.Query(prepare)
	CheckError(err)
	defer stmt.Close()
	for stmt.Next(){
		stmt.Scan(&user.ID)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}