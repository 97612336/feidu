package handlers

import (
	"net/http"
	"feidu/handlers/novel_api"
)

func MyUrls() {
	http.HandleFunc("/test/one", novel_api.Get_some_book_name)
}
