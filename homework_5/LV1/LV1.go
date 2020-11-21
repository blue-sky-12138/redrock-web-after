package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/strong", func(ctx *gin.Context) {
		ctx.String(200,"鑫哥强👍")
	})
	router.Run(":8080")
}