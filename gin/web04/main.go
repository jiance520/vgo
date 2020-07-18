package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter,r *http.Request)  {
	//解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	//是生成的.exe文件的相对路径或绝对路径,go build在哪执行，就生成的目录在哪，打包执行该命令的文件夹
	if err != nil{
		fmt.Printf("HTTP server start failed,err:%v",err)
		return
	}
	//渲染模板
	err=t.Execute(w,"小王子")
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
