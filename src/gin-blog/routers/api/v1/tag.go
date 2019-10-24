package v1

import (
	//"gin-blog/models"
	"../../../models"
	"../../../pkg/e"
	"../../../pkg/util"
	"github.com/astaxie/beego/validation"

	//"gin-blog/pkg/e"
	"../../../pkg/setting"
	//"gin-blog/pkg/setting"
	//"gin-blog/pkg/util"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
//首先获取请求文章名，调用数据处理方法，如果获取到就返回，没有获取到就return
	name :=c.Query("name")
	maps :=make(map[string]interface{})
	data :=make(map[string]interface{})
	if name !=""{
		maps["name"]=name
	}
	var state int = -1
	if age:=c.Query("state");age !=""{
		a,_ := strconv.Atoi(age)
		state = a
		maps["state"]=state
	}
	data["list"] = models.GetTags(util.GetPage(c),setting.PageSize,maps)
	data["total"] =models.GetTagsToal(maps)
	code :=e.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"msg":e.Getmsg(code),
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	name :=c.Query("name")
	a,_ := strconv.Atoi(c.Query("state"))
	state :=a
	createdby :=c.Query("created_by")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdby, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdby, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdby)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.Getmsg(code),
		"data" : make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
	name := c.Query("name")
	b,_ := strconv.Atoi(c.Query("id"))
	id :=b
	a,_ := strconv.Atoi(c.Query("state"))
	state :=a
	ModifiedBy :=c.Query("Modified_by")
	valid := validation.Validation{}
	valid.Required(id, "name").Message("名称不能为空")
	valid.MaxSize(id, 100, "name").Message("名称最长为100字符")
	valid.Required(ModifiedBy, "modified_by").Message("创建人不能为空")
	valid.MaxSize(ModifiedBy, 100, "modified_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		data:= make(map[string]interface{})
		if ! models.ExistTagByID(id) {
			code = e.SUCCESS
			data["name"]=name
			data["state"] = state
			data["modified_by"]=ModifiedBy
			models.EditTag(id,data)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.Getmsg(code),
		"data" : make(map[string]string),
	})


}

//删除文章标签
func DeleteTag(c *gin.Context) {
	b,_ := strconv.Atoi(c.Query("id"))
	id :=b
	a,_ := strconv.Atoi(c.Query("state"))
	state :=a
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(id, "name").Message("名称不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if ! models.ExistTagByID(id) {
			code = e.SUCCESS
			models.DeleatTag(id)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.Getmsg(code),
		"data" : make(map[string]string),
	})
}