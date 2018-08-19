package novel_api

import (
	"net/http"
	"feidu/util"
)

func Show_all_categories(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})

		util.Return_jsonp(w, data)

	}
}
