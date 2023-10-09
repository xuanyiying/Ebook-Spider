package parser

import (
	"ebook-spider/consts"
	"ebook-spider/logger"
	"ebook-spider/model"
	"ebook-spider/store"
	"ebook-spider/utils"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	fetcher "github.com/xuanyiying/fetcher"
	"sort"
	"strconv"
	"strings"
	"time"
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
	// 书籍信息
	main.Find(".ebook-main-wrapper").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h1").Text()                                                                  // 书名
		author := s.Find(".author-name").Text()                                                       // 作者名
		status := strings.TrimPrefix(s.Find(".ebook-status").Text(), "状态：")                           // 状态（完结）
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
	// 图片
	attr, exists := main.Find(".ebook-cover-img").Attr("src")
	if exists {
		r.ImgUrl = attr
		bytes, _ := fetcher.Fetch(attr)
		imgBytes := base64.StdEncoding.EncodeToString(bytes)
		r.ImgBytes = imgBytes
	}
	return r
}

func Run() {
	var books []model.BookInfo
	for i := 8054; i <= 10000 && i >= 0; i-- {
		books = parse(i, books)
	}
	s := store.CsvStore{}
	b, err := s.Store(books)
	if err != nil || !b {
		logger.Error("写入csv错误:", err)
	}
	logger.Info(fmt.Sprintf("成功写入csv:%d 条书籍信息", len(books)))
}

func parse(i int, books []model.BookInfo) []model.BookInfo {
	uri := fmt.Sprintf(consts.AliyunUrl, i)
	logger.Info("url", uri)
	bytes, err := fetcher.FetchWithRateLimiter(uri, time.Tick(time.Second))
	if err != nil {
		logger.Error("错误url:", uri, err)
	}
	a := AliParser{}
	book := a.Parse(bytes, int8(i))
	book.Url = uri
	books = append(books, book)
	sort.Slice(books, func(i, j int) bool {
		t1, _ := time.Parse("2006-01-02", books[i].PublishTime)
		t2, _ := time.Parse("2006-01-02", books[i].PublishTime)
		return t1.Before(t2)
	})
	return books
}
