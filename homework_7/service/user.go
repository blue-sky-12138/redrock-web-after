package service

import (
	"github.com/gin-gonic/gin"
	database "homework_7/database"
	"net/http"
	"strconv"
	"strings"
)

type User struct {					//用户信息结构体
	ID int
	Name string
	Password string
}
type Change struct {				//用于储存用户修改的信息
	ChangeName string
	ChangeNameAgain string
	ChangePassword string
	ChangePasswordAgain string
	ChangeSignature string
}
var (
	change Change
)

//注册
func Register(context *gin.Context)string{
	registerTelephoneNumber,_:= strconv.Atoi(context.PostForm("telephone_number"))
	registerName:=context.PostForm("name")
	registerPassword:=context.PostForm("password")
	registerPasswordAgain:=context.PostForm("passwordAgain")

	//简易判断手机号是否被注册
	name:= database.FindUserName(registerTelephoneNumber)
	if name!=""{
		return "该手机号已注册"
	}

	//简易检查密码
	if strings.Contains(registerPassword,"/"){
		return "密码含有非法字符"
	}else if len(registerPassword)<=6{
		return "密码过短"
	}
	if registerPasswordAgain!=registerPassword{
		return "两次密码不一致"
	}

	registerPassword,registerSalt:= CryptographyNow(registerPassword)
	database.InsertUser(registerName,registerPassword,registerTelephoneNumber,registerSalt)
	return ""
}

//检查用户是否存在
func (user *User)CheckUser()bool{
	OKstring:= database.FindUserName(user.ID)
	if OKstring!=""{
		user.Name=OKstring
		return true
	}
	return false
}

//检查密码是否正确
func (user *User)CheckPassword()bool{
	tem.TemString, tem.TemSalt= database.FindUserPassword(user.Name)
	if Dc.Cryptography(tem.TemSalt); tem.TemString== Dc.Result{
		return true
	}
	return false
}

//更改用户名
func ChangeName(context *gin.Context,ID int,cookie *http.Cookie) bool {
	change.ChangeName=context.PostForm("ChangeName")
	change.ChangeNameAgain=context.PostForm("ChangeNameAgain")
	if change.ChangeName!= change.ChangeNameAgain{
		return false
	}else{
		cookie.Value= change.ChangeName
		database.Update(ID,"name", change.ChangeName)
		return true
	}
}

//更改密码
func ChangePassword(context *gin.Context,ID int)bool{
	change.ChangePassword,_=context.GetPostForm("ChangePassword")
	change.ChangePasswordAgain,_=context.GetPostForm("ChangePasswordAgain")
	if change.ChangePassword!= change.ChangePasswordAgain{
		return false
	}else{
		salt:= database.FindSalt(ID)
		Dc.Data= change.ChangePassword
		Dc.Cryptography(salt)
		database.Update(ID,"password", Dc.Result)
		return true
	}
}

//更改个性签名
func ChangeSignature(context *gin.Context,ID int){
	tem.TemString=context.PostForm("ChangeSignature")
	database.Update(ID,"signature", tem.TemString)
}