package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("chap03-form-bind/web/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", "")
	})

	r.POST("/login", func(c *gin.Context) {
		var form LoginForm

		// 绑定并装配值
		err := c.ShouldBind(&form)
		if err != nil {
			panic(err)
		}

		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Printf("请求参数：%v, [%v - %v]\n", form, username, password)

		c.HTML(http.StatusOK, "index.html", "")
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello World</b>",
		})
	})

	r.GET("/pure-json", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello World</b>",
		})
	})

	r.GET("/get", func(c *gin.Context) {
		username := c.Query("username")
		fmt.Println(username)

		res := []string{"name", "age"}

		// 防止JSON劫持，如果响应数据是数组类型，则默认响应while(1);["name","age"]
		r.SecureJsonPrefix("hh") // 主动设置SecureJSON的头部前缀
		c.SecureJSON(http.StatusOK, res)
	})

	r.Run(":8090")
}
