package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(context *gin.Context) {
		//获取浏览器发送的请求携带的参数 query string
		//name:=context.Query("query") //有多种格式获取前端的参数，如map，数组，query表示前端get方式的key。
		//name:=context.DefaultQuery("query","somebody")//获取胆端的query的键值，如果没有就使用自定义的值somebody
		name, ok := context.GetQuery("query") //如果获取不到的get的参数值，则为false
		if !ok {
			name = "so"
		}
		context.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":9090")
}
