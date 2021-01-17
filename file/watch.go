package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// Watcher 监控文件添加行
type Watcher struct {
	f     *os.File
	r     *bufio.Reader
	close chan bool
}

// New 初始化一个监控程序
func New(filename string) (w *Watcher, oldValue string, err error) {
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}

	// 读入原本的内容
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	oldValue = string(b)

	w = &Watcher{
		f:     f,
		r:     bufio.NewReader(f),
		close: make(chan bool),
	}
	return
}

// Close 结束监控，关闭文件
func (w *Watcher) Close() {
	w.close <- true
	w.f.Close()
}

// Watch 开始进行监控
func (w *Watcher) Watch(callback func(line string)) (err error) {

	//添加要监控的对象，文件或文件夹
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}

	err = watch.Add(w.f.Name())
	if err != nil {
		return
	}

	for {
		select {
		case ev := <-watch.Events:
			if ev.Op&fsnotify.Write == fsnotify.Write {
				// 监控到文件写入
				for {
					line, err := w.ReadLine()
					if err != nil && err != io.EOF {
						return err
					}
					if err == nil {
						// 正确读入一行
						go callback(line)
					} else {
						// 当前是由于 EOF 结束，后面无内容，不需要继续读入下一行
						break
					}
				}
			}
		case err = <-watch.Errors:
			return
		case <-w.close:
			return nil
		}
	}
}

// ReadLine 读入一行的数据
func (w *Watcher) ReadLine() (line string, err error) {
	pos, err := w.f.Seek(0, 1)
	if err != nil {
		return
	}

	line, err = w.r.ReadString('\n')
	if err == io.EOF {
		// 未读到换行符，读到 EOF
		// 放弃本次读取

		// 回退到原本的位置
		w.f.Seek(pos, 0)
		return
	}
	if err != nil {
		// 存在其他错误
		return
	}
	// 清楚末尾的 \n
	line = strings.TrimRight(line, "\n")
	return
}
