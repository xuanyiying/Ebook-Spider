package parser

import (
	"ebook-spider/fetcher"
	"ebook-spider/logger"
	"ebook-spider/model"
	"ebook-spider/store"
	"ebook-spider/utils"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type AliParser struct {
}

func (a *AliParser) Parse(content []byte, id int8) model.BookInfo {
	doc, err := utils.ByteToDocument(content)
	if err != nil {
		logger.Error("获取html失败:", err)
	}
	var r model.BookInfo
	main := doc.Find(".ebook-wrapper")

	main.Find(".ebook-main-wrapper").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h1").Text()                                                                  // 书名
		author := s.Find(".author-name").Text()                                                       // 作者名
		status := s.Find(".ebook-status").Text()                                                      // 状态（完结）
		publishTime := strings.TrimSpace(strings.TrimPrefix(s.Find(".publish-time").Text(), "发布时间：")) // 发布时间（2022-01-12）
		chapterNum := strings.TrimSpace(strings.TrimPrefix(s.Find(".chapter-num").Text(), "章节数："))    // 章节数（16）
		score, _ := strconv.ParseFloat(s.Find(".score-num").Text(), 64)
		stars := score * 0.5
		b := model.BookInfo{
			Id:          id,
			Name:        title,
			Author:      author,
			Status:      status,
			PublishTime: publishTime,
			ChapterNum:  chapterNum,
			Score:       score,
			Stars:       stars,
		}
		r = b
		return
	})
	attr, exists := main.Find(".ebook-cover-img").Attr("src")
	if exists {
		r.ImgUrl = attr
		bytes, _ := fetcher.Fetch(attr)
		imgBytes := base64.StdEncoding.EncodeToString(bytes)
		r.ImgBytes = imgBytes
	}
	return r
}

func Parse() {
	url := "https://developer.aliyun.com/ebook/%d"
	var books []model.BookInfo
	for i := 8046; i < 8051; i++ {
		uri := fmt.Sprintf(url, i)
		logger.Info("url", uri)
		bytes, err := fetcher.Fetch(uri)
		if err != nil {
			logger.Error("错误url:", uri, err)
		}
		a := AliParser{}
		book := a.Parse(bytes, int8(i))
		books = append(books, book)
	}
	s := store.CsvStore{}
	b, err := s.Store(books)
	if err != nil || !b {
		logger.Error("写入csv错误:", err)
	}
	logger.Info(fmt.Sprintf("成功写入csv:%d 条书籍信息", len(books)))
}
