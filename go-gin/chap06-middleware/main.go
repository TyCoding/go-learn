package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()

	r.GET("/hello", helloHandler)

	// 定义认证路由组
	authorized := r.Group("/")
	authorized.Use(auth())
	{
		authorized.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK, "user page")
		})
		authorized.GET("/admin", func(c *gin.Context) {
			c.String(http.StatusOK, "admin page")
		})
	}

	r.Run(":8090")
}

// 同一个组下的请求将首先经过此拦截器
func auth() gin.HandlerFunc {
	fmt.Println("---- auth ----")
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.Contains(path, "admin") {
			fmt.Printf("请求URL=[%v], 拦截此请求\n", path)
			c.String(http.StatusUnauthorized, "Unauthorized page")
			// 拦截请求
			c.Abort()
		}

		// 请求放行
		c.Next()
		return
	}
}

// 自定义接口处理器
func helloHandler(c *gin.Context) {
	println("自定义接口处理器")
}
