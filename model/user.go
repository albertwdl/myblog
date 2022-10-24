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
