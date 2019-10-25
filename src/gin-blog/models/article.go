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
//关联表查找要记得关联
func GetArticle(id int) (articles Article) {
	db.Where("id=?", id).Preload("Tag").First(&articles) //查找数据是否存在
	//db.Where("id=?", id).First(&articles)
	//db.Model(&articles).Related(&articles.Tag) //articles 属于Tag
	return
}

func GetArticlelist(pageNum int,pageSize int,maps interface {})(article []Article){
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&article)
	return
}
func GetArticleTotal(maps interface {}) (count int){
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}
func AddArticlelist(maps map[string]interface{})bool{
	db.Create(&Article{
		TagID:      maps["tag_id"].(int),
		Title:      maps["title"].(string),
		Desc:       maps["desc"].(string),
		Content:    maps["content"].(string),
		CreatedBy:  maps["created_by"].(string),
		State:      maps["state"].(int),
	})
	return true
}


func EditArticle(id int,data interface{})bool{
	db.Model(&Article{}).Where("id= ?", id).Update(data)
	return true
}
func DeleatArticle(id int)bool{
	db.Where("id= ?", id).Delete(&Article{})
	return true
}