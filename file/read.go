package file

/*

	封装了一些文件的操作函数

*/

import (
	"io/ioutil"
	"os"
	"path"

	fp "github.com/OhYee/goutils/functional"
)

// ReadFile 读入文件内容
func ReadFile(filename string) (res []byte, err error) {

	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	res, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}
	return
}

// GetFiles 获取指定目录下的所有文件
func GetFiles(folder string) []string {
	files := make([]string, 0)
	filesInfo, err := ioutil.ReadDir(folder)
	if err == nil {
		for _, f := range filesInfo {
			if !f.IsDir() {
				files = append(files, path.Join(folder, f.Name()))
			}
		}
	}
	return files
}

// GetFolders 获取指定目录下的所有目录
func GetFolders(folder string) []string {
	folders := make([]string, 0)
	filesInfo, err := ioutil.ReadDir(folder)
	if err == nil {
		for _, f := range filesInfo {
			if f.IsDir() {
				folders = append(folders, path.Join(folder, f.Name()))
			}
		}
	}
	return folders
}

// GetFilesWithExt 获取文件夹下具有指定文件后缀得文件列表
func GetFilesWithExt(folder, ext string) []string {
	return fp.FilterString(func(s string, i int) bool {
		return path.Ext(s) == ext
	}, GetFiles(folder))
}
