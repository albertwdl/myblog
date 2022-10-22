package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null;" json:"title"`
	Desc    string `gorm:"type:varchar(200);" json:"desc"`
	Content string `gorm:"type:longtext;" json:"content"`
	Image   string `gorm:"type:varchar(100);" json:"image"`
	Tags    []Tag  `gorm:"many2many:article_tag;" json:"tags"`
}
