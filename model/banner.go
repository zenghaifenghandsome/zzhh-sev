package model

import (
	errorMessages "z-web-sev/utils/errorMessages"

	"gorm.io/gorm"
)
type Banner struct {
	gorm.Model
	Url string
}


//获取所有banner
func GetBanners()([]Banner,int){
	var banners []Banner
	result := db.Find(&banners)

	if result.Error !=nil {
		return nil, errorMessages.ERROR
	}
	return banners,errorMessages.SUCCESS
}