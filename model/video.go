package model

import (
	"fmt"
	errormessages "z-web-sev/utils/errorMessages"

	"gorm.io/gorm"
)

type VideoSource struct {
	gorm.Model
	Key      string `gorm:"type:varchar(100);not null" json:"key"`
	Name     string `gorm:"type:varchar(30);not null" json:"name"`
	Api      string `gorm:"type:varchar(1000);not null" json:"api"`
	Download string `gorm:"type:varchar(1000) default '' " json:"download"`
	ParseUrl string `gorm:"type:varchar(1000) default '' " json:"parseUrl"`
	Group    string `gorm:"type:varchar(20);not null default '默认'" json:"group"`
	IsActive uint   `gorm:"type:int;not null default 1" json:"isActive"`
}

/**
 * get video source list
 */
func GetVideoSourceList() ([]VideoSource, int) {
	var videoSourceList []VideoSource
	result := db.Find(&videoSourceList)
	if result.Error != nil {
		return nil, errormessages.ERROR
	}
	return videoSourceList, errormessages.SUCCESS
}

/**
 * add video source
 */
func AddVideoSource(data *VideoSource) int {
	//fmt.Println("=======================")
	//fmt.Println(data.Key)
	videoSourceExist := IsVideoSourceExist(data.Key)
	if videoSourceExist {
		return errormessages.ERROR
	}
	result := db.Create(&data)
	if result.Error != nil {
		return errormessages.ERROR
	}
	return errormessages.SUCCESS
}

func IsVideoSourceExist(key string) bool {
	var videoSource VideoSource

	result := db.Take(&videoSource).Where("key = ?", key)

	fmt.Println(result.RowsAffected)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

