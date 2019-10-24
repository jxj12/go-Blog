package util

import (
	"../../pkg/setting"
	"github.com/gin-gonic/gin"

	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.Atoi(c.Query("page"))//c.Query() 获取get请求上的键对应的值 com.StrTo()将字符串转换为指定类型
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}