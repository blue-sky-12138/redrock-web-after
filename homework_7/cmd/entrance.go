package cmd

import (
	"github.com/gin-gonic/gin"
	"homework_7/model"
)

var(
	router   *gin.Engine
)

func Entrance(){
	//http://localhost:80/index
	//预载服务器
	router = gin.Default()
	router.LoadHTMLGlob("templates/html/*")
	router.Static("/css", "./templates/css")

	router.GET("/index",model.Index)

	router.GET("/Login",model.LoginGet)
	router.POST("/Login/Create",model.LoginPost)
	router.GET("/Logout",model.LoginOut)

	router.GET("/Register",model.RegisterGet)
	router.POST("/Register/Create",model.RegisterPost)

	router.GET("/ChangeName",model.ChangeNameGet)
	router.POST("/ChangeName/Create",model.ChangeNamePost)

	router.GET("/ChangePassword",model.ChangePasswordGet)
	router.POST("/ChangePassword/Create",model.ChangePasswordPost)

	router.GET("/ChangeSignature",model.ChangeSignatureGet)
	router.POST("/ChangeSignature/Create",model.ChangeSignaturePost)

	router.GET("/MessageWall",model.MessageWall)

	router.GET("/LeaveMessage",model.LeaveMessageGet)
	router.POST("/LeaveMessage/Create",model.LeaveMessagePost)

	router.GET("/Past",model.Past)

	router.GET("/MessageFavorite",model.MessageFavoriteGet)
	router.POST("/MessageFavorite/Create",model.MessageFavoritePost)

	router.GET("MessageDelete",model.MessageDeleteGet)
	router.POST("MessageDelete/Create",model.MessageDeletePost)

	//运行服务器
	router.Run(":8080")
}
