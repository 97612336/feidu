package handlers

import (
	"net/http"
			"weixin_api/handlers/imgapi"
)

func MyUrls() {
	http.HandleFunc("/img/receiver",imgapi.Receive_img)
}
