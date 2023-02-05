package api

import (
	"fmt"
	"net/http"
	"z-web-sev/model"
	errormessages "z-web-sev/utils/errorMessages"
	videoproxy "z-web-sev/utils/videoproxy"

	"github.com/gin-gonic/gin"
)

func GetVideoList(c *gin.Context) {
	videoSourceList, code := model.GetVideoSourceList()
	if code != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errormessages.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errormessages.GetErrMsg(code),
			"data":   videoSourceList,
		})
	}
}

func AddVideoSource(c *gin.Context) {
	var videoSource model.VideoSource

	_ = c.ShouldBindJSON(&videoSource)
	fmt.Print(&videoSource)
	result := model.AddVideoSource(&videoSource)

	if result != errormessages.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": result,
			"mag":    errormessages.GetErrMsg(result),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": result,
			"msg":    errormessages.GetErrMsg(result),
		})
	}

}

func VideProxy(c *gin.Context) {
	url := c.Query("url")
	s, _ := videoproxy.VideoProxys(url)
	c.JSON(http.StatusOK, gin.H{
		"data": s,
	})
}
