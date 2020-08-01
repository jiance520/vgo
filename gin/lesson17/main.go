package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由
func main() {
	r := gin.Default()
	//查看
	r.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	//注册，增加
	r.POST("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	//删除
	r.DELETE("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})
	})
	//更新部分
	r.PUT("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "put",
		})
	})
	//所有的请求方法，使用switch来判断
	r.Any("/user", func(context *gin.Context) {
		switch context.Request.Method {
		case "GET":
			context.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPost:
			context.JSON(http.StatusOK, gin.H{
				"method": "post",
			})
			//...
		}
	})
	//处理用户访问错误的url
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"msg": "noroule",
		})
	})

	//路由组，
	videoGroup := r.Group("/video")
	videoGroup.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "/video/index",
		})
	})
	r.Run(":9090")
}
