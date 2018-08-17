package util

import "net/http"

//获取表单提交的值
func Get_argument(r *http.Request, key string, wantDefault ...string) string {
	argument := r.FormValue(key)
	if argument == "" {
		if wantDefault == nil {
			return ""
		}
		return wantDefault[0]
	}
	return argument
}
