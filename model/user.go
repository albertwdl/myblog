package model

import (
	"encoding/base64"
	"log"
	"myblog/global"
	"myblog/utils/errmsg"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username" binding:"required"`
	Password string `gorm:"type: varchar(20);not null" json:"password" binding:"required"`
	Role     int    `gorm:"type: int" json:"role" binding:"required"`
}

func CheckUser(name string) (uint, int) {
	var user User
	global.DBEngine.Select("id").Where("username = ?", name).Find(&user)
	if user.ID > 0 {
		return user.ID, errmsg.ERROR_USERNAME_USED
	}
	return user.ID, errmsg.SUCCESS
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

func DeleteUser(id uint) int {
	err := global.DBEngine.Delete(&User{}, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户信息
func EditUser(id uint, data *User) int {
	maps := make(map[string]interface{})
	if data.Username != "" {
		maps["username"] = data.Username
	}
	if data.Role != 0 {
		maps["role"] = data.Role
	}
	err := global.DBEngine.Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = ScryptPassword(u.Password)
	return nil
}

func ScryptPassword(password string) string {
	const keyLen = 10
	salt := []byte{82, 47, 96, 90, 12, 38, 45, 92}
	dk, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}

// 登录验证
func CheckLogin(username string, password string) int {
	var user User
	err := global.DBEngine.Where("username = ?", username).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPassword(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
