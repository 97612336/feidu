package novel_api

import (
	"net/http"
	"feidu/util"
	"time"
)

func Get_some_book_name(w http.ResponseWriter, r *http.Request) {
	start_time := time.Now().Unix()
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})
		//获取一组随机数的数组,数组容量为20个
		md5_str := util.Get_md5str("wksgdsg")
		//查询数据库中的数据
		db := util.DB
		rows, err := db.Query("select count(1) from chapter;")
		defer rows.Close()
		util.CheckErr(err)
		var nums_arr []int
		for rows.Next() {
			var count_num int
			err := rows.Scan(&count_num)
			util.CheckErr(err)
			nums_arr = append(nums_arr, count_num)
		}
		data["test_data"] = nums_arr
		data["code"] = 200
		data["msg"] = md5_str
		end_time := time.Now().Unix()
		cost_time := end_time - start_time
		data["time"] = cost_time
		util.Return_json(w, data)
	}
}
