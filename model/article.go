package model

import (
	"myblog/global"
	"myblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null;" json:"title"`
	Desc    string `gorm:"type:varchar(200);" json:"desc"`
	Content string `gorm:"type:longtext;" json:"content"`
	Image   string `gorm:"type:varchar(100);" json:"image"`
	Tags    []Tag  `gorm:"many2many:article_tag;" json:"tags"`
}

func CreateArticle(data *Article) int {
	// 如果没有相应Tag就创建，如果有就设置正确的ID
	for i, v := range data.Tags {
		tagID, code := CheckTag(v.Name)
		if code == errmsg.SUCCESS {
			CreateTag(&v)
			tagID, _ = CheckTag(v.Name)
		}
		data.Tags[i].ID = tagID
	}

	err := global.DBEngine.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetArticles(pageSize int, pageNum int) []Article {
	var articles []Article
	var result *gorm.DB
	if pageSize > 0 && pageNum > 0 {
		result = global.DBEngine.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles)
	} else if pageSize > 0 && pageNum <= 0 {
		result = global.DBEngine.Limit(pageSize).Find(&articles)
	} else {
		result = global.DBEngine.Find(&articles)
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil
	}
	return articles
}

func DeleteArticle(id uint) int {
	err := global.DBEngine.Delete(&Article{}, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditArticle(id uint, data *Article) int {
	for i, v := range data.Tags {
		_, code := CheckTag(v.Name)
		if code == errmsg.SUCCESS {
			CreateTag(&v)
		}
		global.DBEngine.Where("name = ?", v.Name).Find(&data.Tags[i])
	}

	data.ID = id
	err := global.DBEngine.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
