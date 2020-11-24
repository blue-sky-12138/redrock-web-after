package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//http://localhost:8080/cookie
func main() {
	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("abc") // 获取Cookie
		if err != nil {
			cookie = "NotSet"
			// 设置Cookie
			c.SetCookie("abc","password",10,"/","localhost",false,true)
		}
		c.String(http.StatusOK,cookie)
	})

	router.Run()
}