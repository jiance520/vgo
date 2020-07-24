package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(context *gin.Context) {
		//获取浏览器发送的请求携带的参数 query string
		context.HTML(http.StatusOK, "login.html", nil)
	})
	//一次请求对应一次响应
	r.POST("login", func(context *gin.Context) {
		//方式一
		//username:=context.PostForm("username")
		//password:=context.PostForm("password")
		//方式二，如果为空，不使用默认值，如果不存在就使用默认值。
		//username:=context.DefaultPostForm("username","somebody")
		//password:=context.DefaultPostForm("xxx","****")
		//方式三
		username, ok := context.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, ok := context.GetPostForm("password")
		if !ok {
			password = "***"
		}
		context.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9090")
}
