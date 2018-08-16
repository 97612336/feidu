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
