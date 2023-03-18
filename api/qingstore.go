package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qingstor/qingstor-sdk-go/v4/config"
	"github.com/qingstor/qingstor-sdk-go/v4/service"
)

var qy_access_key_id = "PWWJPPEHQKMKKCEOEFAH"
var qy_secret_access_key = "Sm0SUpSXQoARQPhbUmLQrhmd0yJrZgitHC28BNFW"

func UpData(ctx *gin.Context) {
	configuration, _ := config.New(qy_access_key_id, qy_secret_access_key)
	qsService, _ := service.Init(configuration)

	bucket, _ := qsService.Bucket("zzhh-server", "pek3b")

	f, fheader, err := ctx.Request.FormFile("imgfile")
	name := fmt.Sprintf("%d%s", time.Now().Unix(), fheader.Filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result, err2 := bucket.PutObject(name, &service.PutObjectInput{Body: f})
	fmt.Println(service.IntValue(result.StatusCode))
	if err2 != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "上传成功",
		"url":    "https://zzhh-server.pek3b.qingstor.com" + "/" + name,
	})

}
