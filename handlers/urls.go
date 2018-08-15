package handlers

import (
	"net/http"
	"feidu/handlers/imgapi"
	"feidu/handlers/novel_api"
)

func MyUrls() {
	http.HandleFunc("/img/receiver", imgapi.Receive_img)
	http.HandleFunc("/test/one", novel_api.Get_some_book_name)
}
