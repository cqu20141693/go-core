package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// 当前项目根目录
var API_ROOT string

// 查询业务项目的根路径
func GetPath() string {
	// 如果不为空，直接返回
	if API_ROOT != "" {
		return API_ROOT
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		print(err.Error())
	}
	API_ROOT = strings.Replace(dir, "\\", "/", -1)
	return API_ROOT
}

// 判断文件目录是否存在
func IsDirExists(path string) bool {
	// 判断路径的正确性
	stat, err := os.Stat(path)
	if err != nil {
		//
		return os.IsExist(err)
	} else {
		return stat.IsDir()
	}
}
