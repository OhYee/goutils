package file

import (
	"os"
)

// WriteFile 写出文件内容
func WriteFile(filepath string, content []byte) (err error) {
	MkdirForFile(filepath, 0777)
	f, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return
	}

	return
}
