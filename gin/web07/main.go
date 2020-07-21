package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.tmpl")
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
	t.Execute(w, "这是index页面")
}
func home(w http.ResponseWriter, r *http.Request) {
	//先写父模板，再写被包含的子模板
	t, err := template.ParseFiles("home.tmpl")
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
	t.Execute(w, "这是home2页面")
}
func index2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.tmpl", "templates/index2.tmpl") //父模板写前面
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
	t.ExecuteTemplate(w, "index2.tmpl", "这是index2") //要指定渲染哪个
}
func home2(w http.ResponseWriter, r *http.Request) {
	//先写父模板，再写被包含的子模板
	t, err := template.ParseFiles("templates/base.tmpl", "templates/home2.tmpl")
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
	t.ExecuteTemplate(w, "home2.tmpl", "这是home2")
}
func main() { //不要右键执行main,请在命令行执行build+生成的exe
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
}
