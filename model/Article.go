package model

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Tiltle       string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200);not null" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;default:0" json:"read_count"`
}
