package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	database "homework_7/database"
	"strconv"
	"time"
)


type DataCryptographyMD5 struct {	//MD5加密结构体
	Data string
	Result string
}
var (
	Dc     DataCryptographyMD5
	tem    database.Tem
)


//留言
func LeaveMessage(context *gin.Context,name string)bool{
	var (
		message database.Message
		messagePreviousFloor int
		outsideName string
	)
	//处理空字符(用户未输入数据)的情况
	temData:=context.PostForm("PreviousFloor")
	if temData==""{
		messagePreviousFloor=0
	}else {
		messagePreviousFloor,_=strconv.Atoi(temData)
	}

	//判断回复的楼层ID是否存在
	if messagePreviousFloor==0 || database.FindFloorID(messagePreviousFloor) {
		message.Author=name

		//处理是否匿名
		temOutsideName:=context.PostForm("If_anonymity")
		if temOutsideName=="0"{
			outsideName=message.Author
		}else {
			outsideName="Unknown"
		}

		message.Information=context.PostForm("LeaveMessage")
		authorityId,_:=strconv.Atoi(context.PostForm("authorityId"))
		NowTime:=time.Now()
		message.TimeDate=fmt.Sprintf("%d-%d-%d %d:%d:%d",NowTime.Year(),NowTime.Month(),NowTime.Day(),NowTime.Hour(),NowTime.Minute(),NowTime.Second())
		database.InsertMessage(message.TimeDate,message.Author,message.Information,messagePreviousFloor,outsideName,authorityId)
		return true
	}else{
		return false
	}
}

//获取留言墙
func MessageWall(context *gin.Context) *[]database.Message {
	var messages []database.Message
	name,_:=context.Cookie("users")
	memberId:=database.FindUserId(name)
	database.FindAllMessages(&messages,0,0,memberId)
	return &messages
}

//点赞
func MessageFavorite(context *gin.Context) bool {
	floor,_:=strconv.Atoi(context.PostForm("Floor"))
	like,_:=strconv.Atoi(context.PostForm("Like"))
	name,_:=context.Cookie("users")
	if database.FindFloorID(floor){
		database.MessageFavorite(floor,like,name)
		return true
	}else{
		return false
	}
}

//删除留言
func MessageDelete(context *gin.Context,name string) bool {
	floor,_:=strconv.Atoi(context.PostForm("Floor"))
	if database.DeleteQualify(name,floor){
		database.DeleteMessage(floor)
		return true
	}else {
		return false
	}
}