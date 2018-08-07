package adv

import (
	"net/http"
	"fmt"
	"log"
	"weixin_api/models"
	"weixin_api/util"
)

func Add_adver(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.Form.Get("name")
		title := r.Form.Get("title")
		img := r.Form.Get("img")
		media := r.Form.Get("media")
		link := r.Form.Get("link")
		onlinetime := r.Form.Get("onlinetime")
		offlinetime := r.Form.Get("offlinetime")
		kind := r.Form.Get("kind")
		fmt.Println(name, title, img, link, media, onlinetime, offlinetime, kind, nil)
		db := util.Get_sql_db()
		insert_sql := "insert into advert (name,title,img,media,link,kind,onlinetime,offlinetime) values(" +
			"?,?,?,?,?,?,?,?)"
		defer db.Close()
		stmp, prepare_err := db.Prepare(insert_sql)
		if prepare_err != nil {
			fmt.Println(prepare_err)
			return
		}else{
			_, exec_err := stmp.Exec(name, title, img, media, link, kind, onlinetime, offlinetime)
			if exec_err!=nil{
				fmt.Println(exec_err)
				return
			}else{
			//如果运行的都没有错误的话
				var success_json models.Success_json
				success_json=models.Success_json{Code:200,Msg:"添加广告成功"}
				res_str:=util.Get_json_string(success_json)
				fmt.Fprint(w,res_str)
			}
		}
	}
	if r.Method == "GET" {
		r.ParseForm()
		name := r.Form.Get("name")
		fmt.Println(name)
	}

}

func Add(w http.ResponseWriter, r *http.Request)  {
	log.Println("这是为什么呢?")
}