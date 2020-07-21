package main

import (
	"fmt"
	"html/template"
	//"text/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//必须创建模板
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("index.tmpl")
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
	t.Execute(w, "这是index页面")
}
func xss(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err: %v\n", err)
		return
	}
	str1 := "<script>alert(123)</script>"
	str2 := "<a href=li.com>博客</a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}
func main() { //不要右键执行main,请在命令行执行build+生成的exe
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
}
