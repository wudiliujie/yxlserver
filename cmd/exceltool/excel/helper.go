package excel

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if errcode != nil { //忽略错误
		// return errcode
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasPrefix(fi.Name(), "~") {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})
	return files, err
}

func Save(filename string, data string) {
	f, err1 := os.Create(filename) //创建文件
	if err1 != nil {
		fmt.Println(err1)
	}
	_, err1 = io.WriteString(f, data) //写入文件(字符串)
	if err1 != nil {
		fmt.Println(err1)
	}
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
