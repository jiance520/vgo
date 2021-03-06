package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	// 配置静态文件夹路径 第一个参数是自定义任意名的api，第二个是文件夹路径
	r.Static("/xxx", "./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("templates/index.tmpl") //取代template.ParseFiles("templates/base.tmpl"),可加载多个文件
	r.LoadHTMLGlob("templates/**/*")                   //使用正则匹配 所有子目录下的文件。
	r.GET("/posts/index", func(context *gin.Context) { //relativePath任意不重复别名，只是命名不同的接口路径
		//可以用200代替http.StatusOK,
		//gin.H是一个Map，
		//实现原来的包"net/http"的功能
		//函数也可以立时独写在外面，取代http.HandleFunc("/index", index)
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //namel默认为html原文件名，如果使用了{{define "post/index.tmpl"}} 重命名，则必为重新命名的名。
			"title": "posts index.tmpl",
		})
	})
	//路由2
	r.GET("/users/index", func(context *gin.Context) {
		//可以用200代替http.StatusOK,
		//gin.H是一个Map，
		//实现原来的包"net/http"的功能
		//函数也可以立时独写在外面，取代http.HandleFunc("/index", index)
		context.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users index.tmpl <a href='li.com'>博客</a>",
		})
	})
	r.Run(":9090") //启动
}
