package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"z-web-sev/model"
	errormessages "z-web-sev/utils/errorMessages"
	videoproxy "z-web-sev/utils/videoproxy"

	xj "github.com/basgys/goxml2json"
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
	var v interface{}

	url := c.Query("url")
	s, _ := videoproxy.VideoProxys(url)
	xml := strings.NewReader(s)
	jsons, err := xj.Convert(xml)

	json.Unmarshal(jsons.Bytes(), &v)
	data := v.(map[string]interface{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"msg":    "field",
			"data":   "",
		})
	} else {
		//fmt.Println(data)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "success",
			"data":   data,
		})
	}
}

func VideoGet(c *gin.Context) {
	//var v interface{}
	s1 := c.Query("url")
	s2 := c.Query("t")
	s3 := c.Query("pg")
	url := s1 + "?ac=videolist&t=" + s2 + "&pg=" + s3
	s, err := videoproxy.VideoProxys(url)
	//xml := strings.NewReader(s)
	//jsons, err := xj.Convert(xml)

	//json.Unmarshal(jsons.Bytes(), &v)
	//data := v.(map[string]interface{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": s,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": s,
		})
	}

}
