package main

import (
	"log"
	"myblog/global"
	"myblog/model"
	"myblog/routers"
	"myblog/utils/setting"

	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("读取参数失败，请检查参数：%s", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("数据库连接错误：%s", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.AppMode)
	r := routers.NewRouter()
	r.Run(global.ServerSetting.HttpPort)
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Qiniu", &global.QiniuCloudSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
