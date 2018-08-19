package models

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Create_time string `json:"create_time"`
	Url         string `json:"url"`
	Book_img    string `json:"book_img"`
	Kind        string `json:"kind"`
	Author      string `json:"author"`
	Has_chapter string `json:"has_chapter"`
}

type Chapter struct {
	Id           int    `json:"id"`
	Book_id      int    `json:"book_id"`
	Name         string `json:"name"`
	Chapter_text string `json:"chapter_text"`
	Chpater_url  string `json:"chapter_url"`
	Create_time  string `json:"create_time"`
}

type Chapter_id_name struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//banner模型
type Banner_novel struct {
	Book_id int    `json:"book_id"`
	Name    string `json:"name"`
	Author  string `json:"author"`
}

//histroy模型
type View_history struct {
	Book_id    int    `json:"book_id"`
	Chapter_id int    `json:"chapter_id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
}

//hot_novel模型
type Hot_novel struct {
	Book_id int    `json:"book_id"`
	Image   string `json:"images"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
}

//种类模型
type Category struct {
	CategoryId    int    `json:"categoryId"`
	CategoryName  string `json:"categoryName"`
	CategoryImage string `json:"categoryImage"`
}

//类别种类类型
type Category_book struct {
	BookId int    `json:"bookId"`
	Image  string `json:"image"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
}

type One_chapter struct {
	ChapterId int `json:"chapterId"`
	ChapterName string `json:"chapterName"`
	ChapterContent []string `json:"chapterContent"`
}