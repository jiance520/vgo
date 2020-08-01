package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接数据库
	//db,err:=gorm.Open("mysql","root:root@(127.0.0.1:3306)/go?charset=utf8mb4")
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local") //本地时间
	if err != nil {
		panic(err) //导演处理，中止
	}
	defer db.Close()
	//创建表，自动迁移，把结构体和数据表进行对应，增减字结构体字段，数据库表中自动修改。
	db.AutoMigrate(&UserInfo{})
	//创建数据行
	u1 := UserInfo{ID: 1, Name: "七米", Gender: "男", Hobby: "蛙泳"}
	db.Create(u1) //只需要手动创建数据库go，执行go run main.go,数据库表会自动创建user_infos(可以自指定)，并插入数据行，ID相同，不会重复插入，默认把第一个字段做主键
	//查询第一个，并保存到对象,不需要在结构体指定db标识，也不需要像java一样要mapper.xml指定映射
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)
	//更新
	db.Model(&u).Update("hobby", "双色球")
	//删除
	db.Delete(&u)
}
