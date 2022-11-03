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

func GetArticlesByTag(tid uint, pageSize int, pageNum int) ([]Article, int, int) {
	var articles []Article
	var tag Tag
	var total int64

	// 首先检查tagID是否存在
	err := global.DBEngine.Where("id = ?", tid).Find(&tag).Error
	if err != nil {
		return articles, errmsg.ERROR_TAG_NOT_EXIST, 0
	}

	total = global.DBEngine.Model(&tag).Association("Articles").Count()
	if pageSize > 0 && pageNum > 0 {
		err = global.DBEngine.Model(&tag).Limit(pageSize).Offset((pageNum - 1) * pageSize).Association("Articles").Find(&articles)
	} else if pageSize > 0 && pageNum <= 0 {
		err = global.DBEngine.Model(&tag).Limit(pageSize).Association("Articles").Find(&articles)
	} else {
		err = global.DBEngine.Model(&tag).Association("Articles").Find(&articles)
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, int(total)
	}
	return articles, errmsg.SUCCESS, int(total)
}

func GetArticleInfo(id uint) (Article, int) {
	var article Article
	err := global.DBEngine.Preload("Tags").Where("id = ?", id).Take(&article).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return article, errmsg.ERROR_ARTICLE_NOT_EXIST
		} else {
			return article, errmsg.ERROR
		}
	}
	return article, errmsg.SUCCESS
}

func GetArticles(pageSize int, pageNum int) ([]Article, int, int) {
	var articles []Article
	var result *gorm.DB
	var total int64
	if pageSize > 0 && pageNum > 0 {
		result = global.DBEngine.Preload("Tags").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Count(&total)
	} else if pageSize > 0 && pageNum <= 0 {
		result = global.DBEngine.Preload("Tags").Limit(pageSize).Find(&articles).Count(&total)
	} else {
		result = global.DBEngine.Preload("Tags").Find(&articles).Count(&total)
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, int(total)
	}
	return articles, errmsg.SUCCESS, int(total)
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
