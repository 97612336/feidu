package util

import "log"

//检查错误
func CheckErr(err error, args ...string) {
	var hint string
	if len(args) < 1 {
		hint = "Err is: "
	}
	if len(args) >= 1 {
		hint = args[0]
	}
	if err != nil {
		log.Println(hint, err)
	}
}
