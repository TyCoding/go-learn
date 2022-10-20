package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.LoadHTMLGlob("chap04-upload/web/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		if file == nil {
			return
		}

		// 多文件
		//form, _ := c.MultipartForm()
		//files := form.File["upload[]"]

		log.Println(file.Filename)
		dst := "./" + file.Filename
		// 上传到指定路径下
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.Run(":8090")
}
