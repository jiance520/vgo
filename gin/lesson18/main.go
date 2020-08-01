package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//HandlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler")
	c.JSON(http.StatusOK, gin.H{
		"msg": "indexHandler",
	})
}

//定义一个中间件m1，统计程序的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in 中阐件")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort() //可用来没有权限的时候，阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}
func m2(c *gin.Context) {
	fmt.Println("m1 in 中阐件")
	c.Next()
	fmt.Println("m1 out...")
}

func OutMiddlemore(doCheck bool) gin.HandlerFunc {
	return func(context *gin.Context) { //闭包
		if doCheck {
			//业务
		} else {
			context.Next()
		}
	}
}
func main() {
	r := gin.Default()
	r.Use(m1, m2, OutMiddlemore(true)) //全局注册中间件函数m1和m2，使用多个中间件，先进等角和后结束。m1>m2>index>m2>m1
	//GET(relativepath string,handlers ...HandlerFunc) IRoutes //可以传入多个中间件函数。
	//r.GET("/index", m1,indexHandler) //在调用indexHandler函数之前，调用m1，让m1对ndexHandler函数进行耗时统计
	r.GET("/index", indexHandler) //加了r.use(m1)后，这里不要再与
	r.GET("/shop", func(context *gin.Context) {
		fmt.Println("shop in")
		context.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(context *gin.Context) {
		fmt.Println("user in")
		context.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})
	r.Run(":9090")
}
