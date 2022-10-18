package main

import (
	"fmt"
	"myblog/global"
	"myblog/utils/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		fmt.Println("Error is ", err.Error())
	}
}

func main() {
	gin.SetMode(global.ServerSetting.AppMode)
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
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

	return nil
}
