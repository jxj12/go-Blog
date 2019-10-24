package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"` //文章属于标签,相互关联

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (tag *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func GetArticleBYid(id int)bool{
	var articles Article
	db.Select("id").Where("id=?",id).First(&articles) //查找数据是否存在
	if articles.ID >0{
		return true
	}
	return false
}

func GetArticle(id int) (articles Article) {
	db.Where("id=?", id).First(&articles) //查找数据是否存在
	db.Model(&articles).Related(&articles.Tag)
	return
}

func GetArticlelist(pageNum int,pageSize int,maps interface {})(article []Article){
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&article)
	return
}
