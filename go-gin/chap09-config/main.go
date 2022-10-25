package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	//m1()
	m2()
}

// 自定义HTTP配置
func m1() {
	router := gin.Default()
	//http.ListenAndServe(":8090", router)

	s := &http.Server{
		Addr:           ":8090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 >> 20,
	}
	s.ListenAndServe()
}

// 自定义中间件
func m2() {
	r := gin.New() // 不使用gin默认的任何中间件
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println(example)
	})
	r.Run(":8090")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置变量
		c.Set("example", "123")

		// 请求前
		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Println(latency)

		// 获取发送的status
		status := c.Writer.Status()
		log.Println(status)
	}
}
