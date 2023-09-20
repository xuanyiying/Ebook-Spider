package utils

import (
	"bytes"
	"crypto/md5"
	"ebook-spider/consts"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

func MD5(s string) string {
	bytes := md5.Sum([]byte(s))
	r := fmt.Sprintf("%x", bytes) //将[]byte转成16进制
	return r
}

func Str2i(v string) int {
	if i, err := strconv.Atoi(v); err == nil {
		return i
	}
	return 0
}

func GetRuntimeRootPath() string {
	_, f, _, _ := runtime.Caller(0)
	return path.Dir(path.Dir(f))
}

func GetTempDir() string {
	return os.TempDir()
}

func ByteToDocument(content []byte) (*goquery.Document, error) {
	ioReader := bytes.NewReader(content)
	return goquery.NewDocumentFromReader(ioReader)
}

func CsvFilePath() string {
	return GetRuntimeRootPath() + "/" + time.Now().Format("2006-01-02") + consts.CsvExtension
}
