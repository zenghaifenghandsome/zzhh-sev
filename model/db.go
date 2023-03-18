package model

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	dbConfig := viper.GetString("database.DbUser") + ":" + viper.GetString("database.DbPassWord") +
		"@tcp(" + viper.GetString("database.DbHost") + viper.GetString("database.DbPort") + ")/" +
		viper.GetString("database.DbName") + viper.GetString("database.DbProperty")
	//dblogin := "root:zzz000@tcp(localhost:3306)/zeng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	if err != nil {
		fmt.Println(dbConfig)
		fmt.Println("连接数据库失败，请检查参数：", err)
		panic(err)
	} else {
		fmt.Println("sucsess..........................")
		resErr := db.AutoMigrate(
			&User{},
			&UserState{},
			&UserInfo{},
			&BianCheng{},
			&Banner{},
			&BianchengLikes{},
			&Comment{},
			&CommentReply{},
			&Blog{},
			&BlogAuthor{},
			&CheckCode{},
			&VideoSource{},
			&Evd{},
		)
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
