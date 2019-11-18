package api

import (
	"gin-blog/pkg/e"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func  UploadImage(c *gin.Context){
	code := e.SUCCESS
	data :=make(map[string]string)
	file,image,err :=c.Request.FormFile("image")
	if err!=nil{
		logging.Warn(err)
		code =e.ERROR
		c.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.Getmsg(code),
			"data":data,
		})
	}
	if image==nil{
		code =e.INVALID_PARAMS
	}else{
		imageName :=upload.GetImageName(image.Filename)
		fullpath :=upload.GetImageFullPath()
		savePath :=upload.GetImagePath()
		src :=fullpath+imageName
		if ! upload.CheckImageExt(imageName) || upload.CheckImageSize(file){
			code =e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
		}else{
			err :=upload.CheckImage(fullpath)
			if err !=nil{
				logging.Warn(err)
				code=e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			}else if err:= c.SaveUploadedFile(image, src); err != nil{
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			}else{
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.Getmsg(code),
		"data": data,
	})

}