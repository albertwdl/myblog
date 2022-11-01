package model

import (
	"myblog/global"
	"myblog/utils/errmsg"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(20);not null;" json:"name"`
	Articles []Article `gorm:"many2many:article_tag;" json:"articles"`
}

func CheckTag(name string) (uint, int) {
	var tag Tag
	global.DBEngine.Select("id").Where("name = ?", name).Find(&tag)
	if tag.ID > 0 {
		return tag.ID, errmsg.ERROR_TAGNAME_USED
	}
	return tag.ID, errmsg.SUCCESS
}

func CreateTag(data *Tag) int {
	err := global.DBEngine.Create(data)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetTags(pageSize int, pageNum int) []Tag {
	var tags []Tag
	var result *gorm.DB
	if pageSize > 0 && pageNum > 0 {
		result = global.DBEngine.Preload("Articles").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&tags)
	} else if pageSize > 0 && pageNum <= 0 {
		result = global.DBEngine.Preload("Articles").Limit(pageSize).Find(&tags)
	} else {
		result = global.DBEngine.Preload("Articles").Find(&tags)
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil
	}
	return tags
}

func DeleteTag(id uint) int {
	err := global.DBEngine.Delete(&Tag{}, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditTag(id uint, data *Tag) int {
	maps := make(map[string]interface{})
	if data.Name != "" {
		maps["name"] = data.Name
	}
	err := global.DBEngine.Model(&Tag{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
