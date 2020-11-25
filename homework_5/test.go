package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//http://localhost:8080/truck
func main() {
	router := gin.Default()
	router.GET("/truck", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:     "users",
			Value:    "password",
			MaxAge: 1000000,
			Path:     "/",
			Domain: "localhost",
			Secure: false,
			HttpOnly: true,
		}
		http.SetCookie(c.Writer,cookie)
		c.JSON(http.StatusOK, gin.H{"message": cookie})
		c.Request.URL.Path = "/differentWorld"
		router.HandleContext(c)
	})

	router.GET("/differentWorld", Middleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "呐呐呐！你被传送到了异世界哦~"})
	})
	router.Run()
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("users")
		if err == nil {
			value := cookie.Value
			if value == "password" {
				c.JSON(http.StatusOK, gin.H{
					"OK": cookie.Name,
				})
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
			"cookie":cookie,
		})
		c.Abort()
		return
	}
}