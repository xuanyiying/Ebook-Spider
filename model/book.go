package model

type BookInfo struct {
	Id          int8    `csv:"序号"`
	Name        string  `csv:"名字"`
	Author      string  `csv:"作者"`
	Status      string  `csv:"状态"`
	PublishTime string  `csv:"发布时间"`
	ChapterNum  string  `csv:"章节数"`
	Score       float64 `csv:"评分"`
	Stars       float64 `csv:"星"`
	ImgUrl      string  `csv:"图片"`
	ImgBytes    string  `csv:"图片base64数据"`
	Url         string  `csv:"url"`
}
