package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	//浏览器打开http://localhost:9090/user?username=aaa&password=123
	r.GET("/user", func(context *gin.Context) {
		//username:=context.Query("username")
		//password:=context.Query("password")
		//u:=UserInfo{username,password,}
		var u UserInfo
		err := context.ShouldBind(&u) //必须传递地址，才能修改值
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
		context.JSON(http.StatusOK, u)
		fmt.Printf("%#v\n", u)
	})
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})
	//不能浏览器打开，Postman>post body formdata打开http://localhost:9090/form或页面post提交表单
	r.POST("/form", func(context *gin.Context) {
		var u UserInfo                //声明一个UserInfo类型的变量u
		err := context.ShouldBind(&u) //必须传递地址，才能修改值
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
		context.JSON(http.StatusOK, u)
		fmt.Printf("%#v\n", u)
	})
	r.POST("/json", func(context *gin.Context) {
		var u UserInfo                //声明一个UserInfo类型的变量u
		err := context.ShouldBind(&u) //必须传递地址，才能修改值
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
		context.JSON(http.StatusOK, u)
		fmt.Printf("%#v\n", u)
	})
	r.Run(":9090")
}
