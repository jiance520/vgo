package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//浏览器访问地址：http://localhost:9090/王子/18
	r.GET("/:name/:age", func(context *gin.Context) {
		//获取路径参数
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}
