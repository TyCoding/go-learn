package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 写入到文件，每次会覆盖文件原内容
	f, _ := os.Create("gin.log")
	// os.Stdout 日志写入文件同时打印控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	r.Run(":8090")
}
