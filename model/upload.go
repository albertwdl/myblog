package model

import (
	"context"
	"mime/multipart"
	"myblog/global"
	"myblog/utils/errmsg"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: global.QiniuCloudSetting.Bucket,
	}
	mac := qbox.NewMac(global.QiniuCloudSetting.AccessKey, global.QiniuCloudSetting.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := global.QiniuCloudSetting.QiniuServer + ret.Key
	return url, errmsg.SUCCESS
}
