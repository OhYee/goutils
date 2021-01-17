package file

import (
	"os"
	"path"
	"path/filepath"
)

// IsExist 判断文件是否存在
func IsExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		return false

	}
	return true
}

// MkdirForFile 为要访问的文件创建必备的文件夹
func MkdirForFile(filepath string, perm os.FileMode) {
	os.MkdirAll(path.Dir(filepath), perm)
}

// AbsPath 运行目录的绝对路径
func AbsPath() string {
	rpath := path.Join(filepath.Dir(os.Args[0]))
	apath, err := filepath.Abs(rpath)
	if err != nil {
		return rpath
	}
	return apath
}
