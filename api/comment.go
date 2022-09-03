package api

import (
	"net/http"
	"z-web-sev/model"
	errormessages "z-web-sev/utils/errorMessages"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	code := model.AddComment(&data)
	if code != errormessages.SUCCESS {
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"msg":errormessages.GetErrMsg(code),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"msg":errormessages.GetErrMsg(code),
	})
}

func GetComment(c *gin.Context){
	comId := c.Query("id")
	var coms []model.Comment
	coms,code = model.GetComment(comId)
	if code != errormessages.SUCCESS {
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"msg":errormessages.GetErrMsg(code),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"msg":errormessages.GetErrMsg(code),
		"data":coms,
	})
}