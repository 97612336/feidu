package handlers

import (
	"net/http"
	"feidu/handlers/novel_api"
)

func MyUrls() {
	http.HandleFunc("/v1/get_random_book", novel_api.Get_some_book_name)
	http.HandleFunc("/v1/get_chapter_name", novel_api.Get_chapter_name_by_book_id)
	http.HandleFunc("/v1/get_text", novel_api.Get_one_chapter_by_id)
	//小程序主页接口
	http.HandleFunc("/v1/index", novel_api.Index)
	//小程序保存阅读历史的接口
	http.HandleFunc("/v1/save_view_history", novel_api.Save_vie_history)
	//小程序显示所有类别的接口
	http.HandleFunc("/v1/category", novel_api.Show_all_categories)
	//点击种类后显示该种类下的所有书籍
	http.HandleFunc("/v1/category_book", novel_api.Show_category_book)
	//点击书本进入
	http.HandleFunc("/v1/book_read", novel_api.Get_one_chapter_by_book_id)
}
