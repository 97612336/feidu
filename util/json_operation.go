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
	w.Write([]byte(s))
}
