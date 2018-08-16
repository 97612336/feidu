package util

import "net/http"

//获取表单提交的值
func Get_argument(r *http.Request, key string, wantDefault interface{}) interface{} {
	argument := r.FormValue(key)
	if argument == "" {
		if wantDefault == nil {
			return nil
		}
		return wantDefault
	}
	return argument
}
