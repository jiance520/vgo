package main

import (
	"github.com/gin-gonic/gin"
	_ "text/template" //或"html/template"
)

func sayHello(c *gin.Context)  {
	c.JSON(200,gin.H{ //map[string]interface
		"message":"Hello golang", //如果换行，要加逗号,
	})//200是要返回胆端的状态，后面map是返回给前端的值。
}
func main() {
	r :=gin.Default() //返回默认的路由引擎
	//指定用户使用GET请求访问hello时，执行sayHello这个函数
	r.GET("/hello",sayHello)
	//启动服务
	r.Run(":9090")//不写参数，默认是8080，:9090就是简写的localhost:9090
}
