package model

import (
	"myblog/global"
	"myblog/utils/errmsg"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type: varchar(20);not null" json:"username" binding:"required"`
	Password string `gorm:"type: varchar(20);not null" json:"password" binding:"required"`
	Role     int    `gorm:"type: int" json:"role" binding:"required"`
}

func CheckUser(name string) int {
	var users User
	global.DBEngine.Select("id").Where("user_name = ?", name).Find(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

func CreateUser(data *User) int {
	err := global.DBEngine.Create(data)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	var result *gorm.DB
	if pageSize > 0 && pageNum > 0 {
		result = global.DBEngine.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	} else if pageSize > 0 && pageNum <= 0 {
		result = global.DBEngine.Limit(pageSize).Find(&users)
	} else {
		result = global.DBEngine.Find(&users)
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}
