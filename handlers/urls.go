package handlers

import (
	"net/http"
	"feidu/handlers/imgapi"
)

func MyUrls() {
	http.HandleFunc("/img/receiver", imgapi.Receive_img)
}
