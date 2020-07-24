package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(context *gin.Context) {
		//方法一义map
		//data:=map[string]interface{}{
		//	"name":"小王子",
		//	"message":"hello worldt",
		//	"age":18,
		//}
		data := gin.H{
			"name":    "小王子",
			"message": "hello world",
			"age":     18}
		context.JSON(http.StatusOK, data) //对比context.HTML(200， "index.tmpl", gin.H{"title": "博客",})，不用加载模板

	})
	//方法二，返回结构体
	type msg struct {
		Name    string `json:"name"`
		Message string `bson:"message"` //主要被用作MongoDB数据库中的数据存储和网络二进制传输格式
		Age     int
	}
	r.GET("/another_json", func(context *gin.Context) {
		data := msg{Name: "小王子", Message: "hello world", Age: 18}
		context.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
