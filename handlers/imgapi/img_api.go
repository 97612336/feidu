package imgapi

import (
	"net/http"
	"weixin_api/util"
	"log"
	"weixin_api/config"
	"os"
	"io"
)

func Receive_img(w http.ResponseWriter, r *http.Request) {
	//解析表单
	//r.ParseForm()
	r.ParseMultipartForm(1024 * 1024 * 3)
	//判断方法
	if r.Method == "POST" {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		account := util.Get_img_account()
		if account.Upload_name == username && account.Upload_password == password {
			//获取表单文件
			file, header, err := r.FormFile("img")
			if err != nil {
				log.Println("接收表单文件出错")
				log.Println(err)
				return
			}
			defer file.Close()
			//创建写入到本地的文件
			//file_path_and_name := config.Static_Path + header.Filename
			file_path_and_name := config.Static_path_web + header.Filename
			f, err := os.Create(file_path_and_name)
			if err != nil {
				log.Println("创建本地文件的时候出错")
				log.Println(err)
			}
			defer f.Close()
			//url := config.Local_href + header.Filename
			url := config.Static_href + header.Filename
			//执行写入操作
			io.Copy(f, file)
			json_str := "{ \"code\":200," +
				"\"url\":\"" + url + "\"" +
				"}"
			util.Return_json(w, json_str)
			return
		} else {
			json_str := `
			
{
"code":500,
"msg":"输入的账号密码不正确"
}		
`
			util.Return_json(w, json_str)
			return
		}
	} else {
		json_str := `
{
"code":500,
"msg":"请求方式错误"
}
`
		util.Return_json(w, json_str)
		return
	}

}
