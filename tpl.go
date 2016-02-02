package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func LoadTplFiles(dirPth, suffix string) (err error) {
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filepath string, fi os.FileInfo, err error) error {
		//遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			// 开始加载模板
			// !important 此处各文件中tpName会重复，采用用户id+语言+模板名形式区分
			tp := ReadTemplateFile(filepath)
			templates[tp.Uid+tp.Lang+tp.Name] = tp
			// 结束加载模板
		}

		return nil
	})

	return err
}
func ReadTemplateFile(filepath string) Template {
	var tp Template
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))

	//匹配所有路径中的文件名，并切出名称部分
	reg := regexp.MustCompile(`[a-zA-Z]+\.[tp]+`)
	name := reg.FindAllString(filepath, -1)
	name = strings.Split(name[0], ".")
	if name != nil {
		tp.Name = name[0]
	}
	tp.Path = filepath
	tp.Content = string(fd)
	return tp
}
