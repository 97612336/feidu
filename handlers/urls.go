package handlers

import (
	"net/http"
	"feidu/handlers/novel_api"
)

func MyUrls() {
	http.HandleFunc("/v1/get_random_book", novel_api.Get_some_book_name)
	http.HandleFunc("/v1/get_chapter_name", novel_api.Get_chapter_name_by_book_id)
	http.HandleFunc("/v1/get_text", novel_api.Get_one_chapter_by_id)
}
