package routers

import (
	"../pkg/setting"
	v1 "./api/v1"
	"../routers/api"
	"../pkg/upload"
	"github.com/gin-gonic/gin"
	"../middleware/jwt"
	"net/http"
)

func InitRouter() *gin.Engine{
 	r:=gin.New()
 	r.Use(gin.Logger())
 	r.Use(gin.Recovery())
 	//r.Use(jwt.Jwt())

 	gin.SetMode(setting.ServerSetting.RunMode)//SetMode根据输入字符串设置gin模式
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
 	r.GET("/auth",api.GetAuth)
	r.POST("/upload", api.UploadImage)

 	apiv1 :=r.Group("/api/v1")  //路由分组
	apiv1.Use(jwt.Jwt())
 	{
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArtgiticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
 	return r //返回引擎

}
