package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local") //本地时间
	if err != nil {
		panic(err) //导演处理，中止
	}
	defer db.Close()
}
