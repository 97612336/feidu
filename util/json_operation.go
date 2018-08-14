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
func Return_json(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json_str := Get_json_string(i)
	w.Write([]byte(json_str))
}

//返回跨域的json
func Return_jsonp(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	json_str := Get_json_string(i)
	w.Write([]byte(json_str))
}
