package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	b,_ := strconv.Atoi(c.Query("id"))
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

}

//新增文章
func AddArticle(c *gin.Context) {
}

//修改文章
func EditArticle(c *gin.Context) {
}

//删除文章
func DeleteArticle(c *gin.Context) {

}