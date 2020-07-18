package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type User struct {
	Name string
	Gender string
	Age int
}
func sayHello(w http.ResponseWriter,r *http.Request)  {
	//解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	//是生成的.exe文件的相对路径或绝对路径,go build在哪执行，就生成的目录在哪，打包执行该命令的文件夹
	if err != nil{
		fmt.Printf("HTTP server start failed,err:%v",err)
		return
	}
	u1:= User{
		Name:"小王子",
		Gender:"男",
		Age:18,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	m1:=map[string]interface{}{
		"name":"小王子",
		"gender":"男",
		"Age":18,
	}
	//渲染模板
	//err=t.Execute(w,m1)
	//同时传递多个对象，user和 map，使用map传
	err=t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbyList,
	})
	if err != nil{
		fmt.Printf("HTTP server start failed,err:%v",err)
		return
	}
}
func main() {//不要右键执行main,请在命令行执行build+生成的exe
	http.HandleFunc("/",sayHello)
	err :=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Printf("HTTP server start failed,err:%v",err)
	}
	//main.go跟mod在vgo目录下，执行D:\workspace\vgo>go build
	//再执行D:\workspace\vgo>vgo //可省略.exe
}
