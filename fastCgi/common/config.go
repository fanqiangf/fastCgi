package common

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/goless/config"
)

//只开放读的接口

func GetRead(key string) string {

	//获取json配置路径
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file := dir + "/../common/config.json"
	value := getJsonString(key, file)
	return value
}

//读取file文件中的内容
func getJsonString(name string, path string) string {
	config := config.New(path)
	res := config.Get(name)
	ret := getName(res)
	return ret
}

//把拿到的interface接口转换成string
func getName(params ...interface{}) string {

	strArray := make([]string, len(params))
	for i, arg := range params {
		strArray[i] = arg.(string)
	}
	aa := strings.Join(strArray, "_")
	return aa
}
