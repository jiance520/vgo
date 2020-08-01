package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//重定向
func main() {
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		//context.JSON(http.StatusOK,gin.H{
		//	"status":"OK",
		//})
		//自己不返回数据，交给别的接口
		context.Redirect(http.StatusMovedPermanently, "http://www.sogo.com")
	})
	//跟转到其它路由函数，转发?
	r.GET("/a", func(context *gin.Context) {
		context.Request.URL.Path = "/b" //修改请求的URL
		r.HandleContext(context)        //继续后续的处理
	})
	r.GET("/c", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "b",
		})
	})
	r.Run(":9090")
}
