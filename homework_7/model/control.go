package model

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"homework_7/database"
	"homework_7/service"
	"net/http"
	"strconv"
)

var (
	login       service.User
	information   struct{			//主页template信息结构体
		Title string
		VisterOK bool
		LoginOK bool
	}
	PastContent struct{Content string}	//过度template信息结构体
)

func Index(context *gin.Context) {
	tem, err :=context.Request.Cookie("users")
	if err==nil{
		if login.ID==0{
			//获取cookie中的用户数据
			login.Name=tem.Value
			login.ID= database.FindUserTelephone(login.Name)
		}
		information.Title="欢迎回来"+login.Name
		information.LoginOK=true
		information.VisterOK=false
		context.HTML(http.StatusOK,"index.html",information)
	}else{
		information.Title="你好游客"
		information.LoginOK=false
		information.VisterOK=true
		context.HTML(http.StatusOK,"index.html",information)
	}
}
func LoginOut(context *gin.Context) {
	cookie, _ :=context.Request.Cookie("users")
	//删除cookie
	context.SetCookie(cookie.Name,cookie.Value,-1,cookie.Path,cookie.Domain,cookie.Secure,cookie.HttpOnly)
	PastContent.Content="注销成功"
	context.Redirect(http.StatusMovedPermanently,"/Past")
}

//登录界面
func LoginGet(context *gin.Context) {
	context.HTML(http.StatusOK,"login.html",nil)
}
//接受传入登录数据
func LoginPost(context *gin.Context) {
	temName:=context.PostForm("name")
	login.ID,_= strconv.Atoi(temName)
	login.Password=context.PostForm("password")

	//登录检测并重定向
	if login.CheckUser() && login.CheckPassword(){
		cookie := &http.Cookie{
			Name:     "users",
			Value:    login.Name,
			MaxAge: 100000,
			Path:     "/",
			Domain: "localhost",
			Secure: false,
			HttpOnly: true,
		}
		http.SetCookie(context.Writer,cookie)
		context.Redirect(http.StatusMovedPermanently,"/index")
	}else{
		PastContent.Content="请检查你的用户名和密码是否正确"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}

//注册界面
func RegisterGet(context *gin.Context) {
	context.HTML(http.StatusOK,"register.html",nil)
}
func RegisterPost(context *gin.Context) {
	err:= service.Register(context)
	switch {
	case err=="":
		PastContent.Content="注册成功"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	default:
		PastContent.Content=err
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}

//更改名称
func ChangeNameGet(context *gin.Context) {
	context.HTML(http.StatusOK,"ChangeName.html",nil)
}
func ChangeNamePost(context *gin.Context) {
	cookie,_:=context.Request.Cookie("users")
	OK:= service.ChangeName(context,login.ID,cookie)
	if OK{
		PastContent.Content="更改完成"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}else{
		PastContent.Content="两次输入的名称不一致！"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}

//更改密码
func ChangePasswordGet(context *gin.Context) {
	context.HTML(http.StatusOK,"ChangePassword.html",nil)
}
func ChangePasswordPost(context *gin.Context) {
	OK:= service.ChangePassword(context,login.ID)
	if OK{
		PastContent.Content="两次输入的密码不一致！"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}else{
		PastContent.Content="更改完成"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}

//更改个性签名
func ChangeSignatureGet(context *gin.Context) {
	context.HTML(http.StatusOK,"ChangeSignature.html",nil)
}
func ChangeSignaturePost(context *gin.Context) {
	service.ChangeSignature(context,login.ID)
	PastContent.Content="更改完成"
	context.Redirect(http.StatusMovedPermanently,"/Past")
}

//留言墙
func MessageWall(context *gin.Context) {
	message:= service.MessageWall(context)
	context.HTML(http.StatusOK,"MessageHistory.html",message)
}

//留言
func LeaveMessageGet(context *gin.Context) {
	context.HTML(http.StatusOK,"LeaveMessage.html",nil)
}
func LeaveMessagePost(context *gin.Context) {
	OK:= service.LeaveMessage(context,login.Name)
	if OK{
		PastContent.Content="留言成功"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}else{
		PastContent.Content="回复的楼层的ID不存在！"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}

//过渡页面(3秒后回到主页)
func Past(context *gin.Context) {
	context.HTML(http.StatusOK,"Past.html",PastContent)
}

//点赞
func MessageFavoriteGet(context *gin.Context) {
	context.HTML(http.StatusOK,"MessageFavorite.html",nil)
}
func MessageFavoritePost(context *gin.Context) {
	OK:=service.MessageFavorite(context)
	if OK{
		PastContent.Content="操作完成"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}else{
		PastContent.Content="回复楼层不存在"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}

//删除留言
func MessageDeleteGet(context *gin.Context) {
	context.HTML(http.StatusOK,"MessageDelete",nil)
}
func MessageDeletePost(context *gin.Context) {
	cookieName,_:=context.Cookie("users")
	OK:=service.MessageDelete(context,cookieName)
	if OK{
		PastContent.Content="删除完成"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}else {
		PastContent.Content="该楼层ID不存在或不是您的留言"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}