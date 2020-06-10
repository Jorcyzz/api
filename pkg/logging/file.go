package logging

import (
	"api/pkg/file"
	"api/pkg/setting"
	"fmt"
	"os"
	"time"
)

//获取日志文件路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

//获取日志文件的名称
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

//打开日志文件
func openLogFile(fileName, filePath string) (*os.File, error) {

	dir, err := os.Getwd() //返回与当前目录对应的根路径名
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src) //检查文件权限
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src) //如果不存在则新建文件夹
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	//调用文件，在写入时将数据追加到文件中 | 如果不存在，则创建一个新文件 | 以只写模式打开文件
	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
