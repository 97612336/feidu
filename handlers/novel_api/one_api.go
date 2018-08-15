package novel_api

import (
	"net/http"
	"feidu/util"
)

func Get_some_book_name(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})
		//获取一组随机数的数组,数组容量为20个
		md5_str := util.Get_md5str("wksgdsg")
		data["code"] = 200
		data["msg"] = md5_str
		util.Return_json(w,data)
	}
}
