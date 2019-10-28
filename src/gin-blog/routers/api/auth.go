package api
import "C"
import (
	"../../models"
	"../../pkg/e"
	"../../pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuth(c *gin.Context){
	username :=c.Query("username")
	password :=c.Query("password")

	valid :=validation.Validation{}
	valid.MaxSize(username, 50, "username").Message("最长为50字符")
	valid.MaxSize(password, 50, "password").Message("最长为50字符")
	data :=make(map[string]interface{})
	code :=e.INVALID_PARAMS
	isExist := models.CheckAuth(username, password)//查询用户
	if isExist{
		token,err :=util.GenerateToken(username,password)
		if err !=nil{
			code =e.ERROR_AUTH_TOKEN
		}else {
			code =e.SUCCESS
			data["token"]=token
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"msg":e.Getmsg(code),
	})

}