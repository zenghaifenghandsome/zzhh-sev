package api

import (
	"net/http"
	"z-web-sev/model"
	errormessages "z-web-sev/utils/errorMessages"

	"github.com/gin-gonic/gin"
)

func GetBanners(c *gin.Context) {
	banners, err := model.GetBanners()
	if err != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"msg": "获取轮播图失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": err,
		"msg":  errormessages.GetErrMsg(err),
		"data": banners,
	})
}
