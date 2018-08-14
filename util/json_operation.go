package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//将类型转化为字符串json
func Get_json_string(m interface{}) string {
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	res_str := string(res)
	return res_str
}

//在web中返回json字符串
func Return_json(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(s))
}

//直接根据类型对象返回字符串响应
func Write_json(w http.ResponseWriter, i interface{}) {
	json_str := Get_json_string(i)
	w.Write([]byte(json_str))
}
