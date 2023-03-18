package utils

import (
	"context"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func UpDate() {
	baseUrl, _ := url.Parse("https://zengdd-1306364512.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BatchURL: baseUrl}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDlZByYu40lMkwwlOp26xew0rVwBFE6oNk",
			SecretKey: "GZiwysdJrGVYP2VnuoX8mPq6B5mFbmmZ",
		},
	})
	_, error := c.Object.PutFromFile(context.Background(), "name", "../", nil)
	if error != nil {
		panic(error)
	}
}
