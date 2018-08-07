package index

import (
	"net/http"
	"io/ioutil"
	"html/template"
	"os"
	"weixin_api/config"
	"weixin_api/models"
	"weixin_api/util"
	"io"
	"log"
)

//主页面跳转
func Index(w http.ResponseWriter, r *http.Request) {
	//读取目录下的所有文件
	static_dir := config.Static_Path
	files, _ := ioutil.ReadDir(static_dir)
	//定义所有文件组成的数组
	var file_list []map[string]string
	for _, v := range files {
		var one_file = make(map[string]string)
		one_file["name"] = v.Name()
		one_file["href"] = config.Static_href + v.Name()
		file_list = append(file_list, one_file)
	}
	//渲染模板
	tmpl, _ := template.ParseFiles("template/index.html")
	//创建结果字典
	var Res models.Index
	Res.Files = file_list
	tmpl.Execute(w, Res)
	return
}

//去往上传文件页,并执行验证
func Add_file(w http.ResponseWriter, r *http.Request) {
	//获取表单提交的信息
	r.ParseForm()
	if r.Method == "POST" {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		info := util.Get_img_account()
		if info.Upload_name == username {
			if info.Upload_password == password {
				tmpl, _ := template.ParseFiles("template/upload_file.html")
				tmpl.Execute(w, "")
				return
			} else {
				tmpl, _ := template.ParseFiles("template/err.html")
				tmpl.Execute(w, "")
				return
			}
		} else {
			tmpl, _ := template.ParseFiles("template/err.html")
			tmpl.Execute(w, "")
			return
		}
	}

}

//具体的添加文件的操作
func Upload_file(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var err error
	//获取表单文件
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("接收表单文件出错")
		log.Println(err)
		return
	}
	defer file.Close()
	//创建写入到本地的文件
	file_path_and_name := config.Static_Path + header.Filename
	f, err := os.Create(file_path_and_name)
	if err != nil {
		log.Println("创建本地文件的时候出错")
		log.Println(err)
	}
	defer f.Close()
	//执行写入操作
	io.Copy(f, file)
	//进行页面跳转
	http.Redirect(w, r, "/", http.StatusFound)
	return
}
