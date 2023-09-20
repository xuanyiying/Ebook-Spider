package store

import (
	"ebook-spider/logger"
	"ebook-spider/utils"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"reflect"
)

type CsvStore struct{}

func (c *CsvStore) Store(data any) (bool, error) {
	path := utils.CsvFilePath()
	logger.Info("path:", path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		logger.Error("err:", err)
		return false, err
	}
	defer file.Close()
	err = gocsv.MarshalFile(data, file)
	if err != nil {
		logger.Error("写入csv错误:", err)
		return false, err
	}
	logger.Info(fmt.Sprintf("成功写入csv:%d 条数据", reflect.ValueOf(data).Len()))
	return true, nil
}
