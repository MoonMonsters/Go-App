package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

/**
获取文件大小
*/
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

/**
获取文件名的后缀
*/
func GetExt(filename string) string {
	return path.Ext(filename)
}

/**
检查文件是否存在
1. 如果返回nil, 则表示文件或目录存在
2. 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
3. 如果返回的错误为其它类型,则不确定是否在存在
*/
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return false
}

/**
检查文件权限
*/
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

/**
新建目录
*/
func MkDir(src string) error {
	// 第二个参数是目录权限
	err := os.MkdirAll(src, os.ModePerm)

	return err
}

/**
判断目录是否存在, 不存在则新建
*/
func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

/**
使用指定模式打开文件
*/
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, err
}
