package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	k := func(name string) (string, error) {
		return name + "年轻", nil
	}
	//t,err:=template.New("f.tmpl").ParseFiles("f.tmpl")
	t := template.New("f.tmpl") //必须跟后面的parsefile其的一个文件名一样
	//在解析模板前，把自定义函数注册到模板中。
	t.Funcs(template.FuncMap{
		"kuo": k,
	})
	//解析模板
	_, err := t.ParseFiles("./f.tmpl")
	//是生成的.exe文件的相对路径或绝对路径,go build在哪执行，就生成的目录在哪，打包执行该命令的文件夹
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
	t.Execute(w, "小王子")
}
func demo1(w http.ResponseWriter, r *http.Request) {
	//先写父模板，再写被包含的子模板
	t, err := template.ParseFiles("t.tmpl", "ul.tmpl")
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
	t.Execute(w, "小王子")
}
func main() { //不要右键执行main,请在命令行执行build+生成的exe
	http.HandleFunc("/", f1)
	http.HandleFunc("/template", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
	//main.go跟mod在vgo目录下，执行D:\workspace\vgo>go build
	//再执行D:\workspace\vgo>vgo //可省略.exe
}
