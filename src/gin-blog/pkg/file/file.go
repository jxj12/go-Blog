package file

import (
	"io/ioutil"
	"mime/multipart" //实现了 MIME 的 multipart 解析
	"path"
	"os"
)
//获取文件大小
func GitSize(f multipart.File)(int ,error){
	file,err :=ioutil.ReadAll(f)
	return len(file),err
}
//获取文件后缀
func GitExt(filename string)(string){
	return path.Ext(filename) //返回路径中扩展部分,获取文件类型
}
//检查文件是否存在
func CheckExist(src string)bool{ //校验
	_,err:=os.Stat(src)
	return os.IsNotExist(err)//判定err错误是否是权限错
}

func MkDir(src string)error{
	err :=os.MkdirAll(src,os.ModePerm)//os.ModePerm永久文件模式
	if err !=nil{
		return err
	}
	return nil

}
//如果不存在则新建文件夹
func IsNotExistMkDir(src string)error{ //文件不存在
	if exist :=CheckExist(src);exist == false{
		if err :=MkDir(src);err !=nil{
			return err
		}
	}
	return nil
}
//检查文件权限
func CheckPermission(src string)bool{
	_,err:=os.Stat(src)
	return os.IsPermission(err)//判定err错误是否是权限错

}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}



