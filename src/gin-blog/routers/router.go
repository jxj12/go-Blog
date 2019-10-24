package routers

import (
	"../pkg/setting"
	v1 "./api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
 	r:=gin.New()
 	r.Use(gin.Logger())
 	r.Use(gin.Recovery())
 	gin.SetMode(setting.RunMode)//SetMode根据输入字符串设置gin模式
 	apiv1 :=r.Group("/api/v1")  //路由分组
 	{
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

	}
 	return r //返回引擎
}
