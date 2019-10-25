package v1

import (
	"../../../models"
	"../../../pkg/e"
	"../../../pkg/util"
	"../../../pkg/setting"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	b,_ := strconv.Atoi(c.Param("id"))
	id :=b
	valid := validation.Validation{}
	valid.Required(id, "id").Message("名称不能为空")
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors(){
		if !models.GetArticleBYid(id){
			code =e.ERROR_EXIST_TAG
		}
		code =e.SUCCESS
		data=models.GetArticle(id)
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":e.Getmsg(code),
		"data" : data,
	})
}

//获取多个文章
func GetArticles(c *gin.Context) {
	data :=make(map[string]interface{})
	maps :=make(map[string]interface{})
	id,_:=strconv.Atoi(c.Query("tag_id"))
	s ,_:= strconv.Atoi(c.Query("state"))
	state:=s
	tagId :=id
	valid:=	validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code :=e.INVALID_PARAMS
	if !valid.HasErrors(){
		code =e.SUCCESS
		data["list"]=models.GetArticlelist(util.GetPage(c),setting.PageSize,maps)
		data["Total"]=models.GetArticleTotal(maps)
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.Getmsg(code),
		"data" : data,
	})

}

//新增文章
func AddArticle(c *gin.Context) {
	id,_:=strconv.Atoi(c.Query("tag_id"))
	tagId :=id
	title :=c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	s ,_:= strconv.Atoi(c.Query("state"))
	state:=s
	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code :=e.INVALID_PARAMS
	if !valid.HasErrors(){
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"]=state
			models.AddArticlelist(data)
			code =e.SUCCESS
		}else{
			code=e.ERROR_NOT_EXIST_TAG
		}
	}else{
		for _,err :=range valid.Errors{
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.Getmsg(code),
		"data":make(map[string]interface{}),
	})

}

//修改文章
func EditArticle(c *gin.Context) {
	d, _ := strconv.Atoi(c.Param("id"))
	id := d
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	s, _ := strconv.Atoi(c.Query("state"))
	state := s
	i, _ := strconv.Atoi(c.Query("tag_id"))
	tagId := i
	data := make(map[string]interface{})
	data["state"] = state
	data["desc"] = desc
	data["content"] = content
	data["modifiedBy"] = modifiedBy
	data["tag_id "]=tagId
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	code := e.SUCCESS
	if !valid.HasErrors() {
		if models.GetArticleBYid(id) {
			data := make(map[string]interface{})
			data["state"] = state
			data["desc"] = desc
			data["content"] = content
			data["modifiedBy"] = modifiedBy
			data["tag_id "]=tagId
			data["title"]=title
			models.EditArticle(id, data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.Getmsg(code),
		"data":data,
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	d, _ := strconv.Atoi(c.Param("id"))
	id := d
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code :=e.INVALID_PARAMS
	if !valid.HasErrors(){
		if models.GetArticleBYid(id){
			code =e.SUCCESS
			models.DeleatArticle(id)
		}else{
			code =e.ERROR
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.Getmsg(code),
		"data":make(map[string]interface{}),
	})
}