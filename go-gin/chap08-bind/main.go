package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type FormA struct {
	// 注意，第一个名称必须是大写开头，否则绑定不上
	Name string `form:"name" binding:"required"`
	// 注意：这个ID要全大写才能绑定上
	ID  int    `uri:"id"`
	Tag string `uri:"tag"`
}
type FormB struct {
	Name string `form:"name" binding:"required"`
}

type Form struct {
	Des   string `form:"des"`
	FormA *FormA
	Group struct {
		ID   int    `form:"id"`
		Name string `form:"name"`
	}
}

func main() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		// request body的多次绑定
		var objA FormA
		var objB FormB
		err := c.ShouldBind(&objA)
		if err != nil {
			panic(err)
		}
		fmt.Printf("formA绑定的参数：%v\n", objA)

		// c.ShouldBind只能调用一次，第二次绑定会报错
		//err2 := c.ShouldBind(&objB)
		//if err2 != nil {
		//	panic(err2)
		//}

		// 想要绑定多次表单，可以使用c.ShouldBindBodyWith()
		//c.ShouldBindBodyWith(&objA, binding.JSON)
		c.ShouldBindBodyWith(&objB, binding.JSON)

		fmt.Printf("formB绑定的参数：%v\n", objB)

		// 127.0.0.1:8090/form?id=11&ids[1]=a&ids[2]=b
		fmt.Printf("查询url中的数据： id=%v, ids=%v\n", c.Query("id"), c.QueryMap("ids"))
		fmt.Printf("查询body中的数据： form=%v\n", c.PostFormMap("name"))
	})

	// 常见的，一般是绑定json格式数据
	r.POST("/form2", func(c *gin.Context) {
		var form FormA
		c.ShouldBindJSON(&form)
		fmt.Printf("formA绑定参数：%v\n", form)
	})

	// 绑定url中的参数
	r.GET("/form3/:id/:tag", func(c *gin.Context) {
		var form FormA
		c.ShouldBindUri(&form)
		fmt.Printf("formA绑定参数：%v\n", form)
	})

	// 绑定复杂结构体
	r.POST("/form4/:id/:tag", func(c *gin.Context) {
		var form Form
		c.ShouldBind(&form)
		fmt.Printf("formA绑定参数：%v\n", form)
	})

	r.Run(":8090")
}
