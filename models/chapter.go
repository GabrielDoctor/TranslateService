package models

type Chapter struct {
	NovelId      int
	VolumeId     int
	ChapterId    int
	ChapterTitle string
	ContentPath  string
	Content      string
	UpTime       string
	AuthorSay    string
}
