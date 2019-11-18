package upload

import (
	"../setting"
	"fmt"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"../file"
)

//获取图片完整访问URL
func GetImageFullUrl(name string)string{
	return setting.AppSetting.ImagePrefixUrl+"/"+GetImagePath()+name
}

//获取jia mi图片名称
func GetImageName(name string)string{
	ext :=path.Ext(name)
	filename :=strings.TrimSuffix(name,ext) //name字符串如果以ext结尾则去掉，否则返回原始字符串name
	filename =util.EncodeMD5(filename)
	return filename + ext
}

//获取图片路径
func GetImagePath()string{
	return setting.AppSetting.ImageSavePath
}

//获取图片完整路径
func GetImageFullPath()string{
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

//检查图片后缀
func CheckImageExt(fileName string)bool{
	ext := file.GitExt(fileName)
	for _,allowExt :=range setting.AppSetting.ImageAllowExts{
		if strings.ToUpper(allowExt)==strings.ToUpper(ext){
			return true
		}
	}
	return false
}

//检查图片大小
func CheckImageSize(f multipart.File)bool{
	size,err:=file.GitSize(f)
	if err !=nil{
		log.Println(err)
		logging.Warn(err)
		return false
	}
	return size <=setting.AppSetting.ImageMaxSize
}

//检查图片
func CheckImage(src string)error{
	dir,err:=os.Getwd()
	if err!=nil{
		return fmt.Errorf("os.Getwd err:%v",err)
	}
	err =file.IsNotExistMkDir(dir+"/"+src)
	if err !=nil{
		return fmt.Errorf("file.IsNotExistMkDir err:%v",err)
	}
	perm :=file.CheckPermission(src)
	if perm == true{
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}
