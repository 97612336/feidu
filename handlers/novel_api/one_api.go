package novel_api

import "net/http"

func Get_some_book_name(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		//获取一组随机数的数组,数组容量为20个

	}
}
