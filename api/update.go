package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func UpDate(ctx *gin.Context) {
	baseUrl := "https://zengdd-1306364512.cos.ap-shanghai.myqcloud.com"
	u, _ := url.Parse(baseUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDlZByYu40lMkwwlOp26xew0rVwBFE6oNk",
			SecretKey: "GZiwysdJrGVYP2VnuoX8mPq6B5mFbmmZ",
		},
	})

	//name := "test/objectput.jpg"

	//f, err := os.Open("C://Users//zengh//Desktop//5c3c751112b7a119a4d696419677cbf.jpg")
	f, fheader, err := ctx.Request.FormFile("imgfile")
	name := fmt.Sprintf("%d%s", time.Now().Unix(), fheader.Filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err2 := c.Object.Put(context.Background(), name, f, nil)
	if err2 != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "上传成功",
		"url":    baseUrl + "/" + name,
	})
}
