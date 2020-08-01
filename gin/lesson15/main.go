package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//处理multiport forms提交文件时默认内存限制是32MB
	//r.MaxMultiportMemory=8<<20 //修改大小。8MB
	r.LoadHTMLFiles("index.html")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/upload", func(context *gin.Context) {
		f, err := context.FormFile("f1") //获取文件
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//将文件保存在服务器上。
			dst := fmt.Sprintf("./%s", f.Filename) // 拼 接保存的路径和文件名
			//dst:=path.Join("./",f.Filename) //可 接收多个参数
			context.SaveUploadedFile(f, dst)
			context.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}

		//多文件上传
		//form,_:=context.MultipartForm()
		//files:= form.File["file"]
		//for index,file:=range files{
		//	log.Println(file.Filename)
		//	dst:=fmt.Sprintf("c:/tmp/%s_%d",file.Filename,index)
		//	context.SaveUploadedFile(file,dst)
		//}
		//context.JSON(http.StatusOK,gin.H{
		//	"message":fmt.Sprintf("%d files uploaded!",len(files)),
		//})
	})
	r.Run(":9090")
}
