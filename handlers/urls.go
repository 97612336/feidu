package handlers

import (
	"net/http"
	"feidu/handlers/novel_api"
)

func MyUrls() {
	http.HandleFunc("/test/one", novel_api.Get_some_book_name)
	http.HandleFunc("/test/two", novel_api.Get_chapter_name_by_book_id)
	http.HandleFunc("/test/three", novel_api.Get_one_chapter_by_id)
}
