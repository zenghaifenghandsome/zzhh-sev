package api

import (
	"context"
	"fmt"
	"net/http"
	errormessages "z-web-sev/utils/errorMessages"

	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var ImageUrl = "http://qiniu.zzhh.asia/"
var AccessKey = "cztyOFu4ykkkrD4ubQwFAyZkR2gw5WwufdbmVlYM"
var SecretKey = "MzcQrh0tQRZ1iNj7WSDm72nCLRiFOueApcIj9QL6"
var Bucket = "zenghaifeng"

func UpDate(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("imgfile")

	fileSize := fileHeader.Size
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	imgUrl, err := UpLodeFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status": err,
		"msg":    errormessages.GetErrMsg(err),
		"imgUrl": imgUrl,
	})
}

func UpLodeFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	PutExtra := storage.PutExtra{}

	formUpaloader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}

	err := formUpaloader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &PutExtra)
	if err != nil {
		return "", errormessages.ERROR
	}
	url := ImageUrl + ret.Key
	return url, errormessages.SUCCESS

}
