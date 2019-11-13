package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name string `json:"name"` //'标签名称'
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"` //'状态 0为禁用、1为启用'
}



func GetTags(pageNum int,pageSize int,maps interface {})(tags []Tag){
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagsToal(maps interface{})(count int){

	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string)bool{
	var tag Tag
	db.Select("id").Where("name=?",name).First(&tag) //查找数据是否存在
	if tag.ID >0{
		return true
	}
	return false
}

func AddTag(name string,state int,createdBy string)bool{
	//如果不存在就添加
	db.Create(&Tag{
		Name:       name,
		CreatedBy:  createdBy,
		State:      state,
	})
	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix()) //设置列的值
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func ExistTagByID(id int)bool{
	var tags Tag
	db.Select("id").Where("id=?",id).First(&tags) //查找数据是否存在
	if tags.ID >0{
		return true
	}
	return false
}

func EditTag(id int,data interface{})bool{
	db.Model(&Tag{}).Where("id= ?", id).Update(data)
	return true
}

func DeleatTag(id int)bool{
	db.Where("id= ?", id).Delete(&Tag{})
	return true
}

func ClentAllTage()bool{
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})
	return true
}