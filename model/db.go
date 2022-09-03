package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	dblogin := "root:niganshenmelo1@tcp(101.35.101.200:3306)/zzhh?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dblogin), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		panic(err)
	} else {
		fmt.Println("sucsess..........................")
		resErr := db.AutoMigrate(&User{}, &UserState{}, &UserInfo{}, &BianCheng{}, &Banner{}, &BianchengLikes{}, &Comment{}, &CommentReply{}, &Blog{}, &BlogAuthor{}, &CheckCode{})
		if resErr != nil {
			fmt.Println(resErr)
		}
		// db.AutoMigrate(&UserInfo{})
		// db.AutoMigrate(&BianCheng{})
		// db.AutoMigrate(&Comment{})

		sqlDB, _ := db.DB()

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxIdleTime(10 * time.Hour)

	}
}
