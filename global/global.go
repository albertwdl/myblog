package global

import (
	"myblog/utils/setting"

	"gorm.io/gorm"
)

var (
	ServerSetting     *setting.ServerSettingS
	DatabaseSetting   *setting.DatabaseSettingS
	DBEngine          *gorm.DB
	QiniuCloudSetting *setting.QiniuCloudSettingS
)
