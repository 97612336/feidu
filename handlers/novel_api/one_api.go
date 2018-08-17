package novel_api

import (
	"net/http"
	"feidu/util"
	"feidu/models"
	"strconv"
)

func Get_some_book_name(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})
		var books []models.Book
		sql_str := "select * from book where has_chapter=1 order by rand() limit 10;"
		rows, err := util.DB.Query(sql_str)
		util.CheckErr(err)
		for rows.Next() {
			var one_book models.Book
			err := rows.Scan(&one_book.Id, &one_book.Name, &one_book.Create_time, &one_book.Url,
				&one_book.Book_img, &one_book.Kind, &one_book.Author, &one_book.Has_chapter)
			util.CheckErr(err)
			books = append(books, one_book)
		}
		data["books"] = books
		data["code"] = 200
		util.Return_jsonp(w, data)
	}
}

func Get_chapter_name_by_book_id(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})
		var id_names []models.Chapter_id_name
		book_id := r.FormValue("book_id")
		n := r.FormValue("n")
		page, err := strconv.Atoi(n)
		util.CheckErr(err)
		tmp_page_size := r.FormValue("page_size")
		page_size, err := strconv.Atoi(tmp_page_size)
		util.CheckErr(err)
		sql_str := "select id,name from chapter where book_id =" + book_id + " limit " + strconv.Itoa((page-1)*page_size) +
			"," + strconv.Itoa(page_size) + ";"
		rows, err := util.DB.Query(sql_str)
		util.CheckErr(err)
		for rows.Next() {
			var one models.Chapter_id_name
			err := rows.Scan(&one.Id, &one.Name)
			util.CheckErr(err)
			id_names = append(id_names, one)
		}
		data["names"] = id_names
		data["code"] = 200
		util.Return_jsonp(w, data)
	} else if r.Method == "POST" {

	}
}

func Get_one_chapter_by_id(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})
		chapter_id := r.FormValue("chapter_id")
		sql_str := "select chapter_text from chapter where id =" + chapter_id + ";"
		rows, err := util.DB.Query(sql_str)
		util.CheckErr(err)
		var one_text string
		for rows.Next() {
			rows.Scan(&one_text)
		}
		//把字符串转化为字符串组成的数组
		var res_text []string
		util.Json_to_object(one_text, &res_text)
		data["code"] = 200
		data["text"] = res_text
		util.Return_jsonp(w, data)
	}
}

//首页接口
func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		user_id := r.FormValue("user_id")
		var data = make(map[string]interface{})
		//首先获取banner
		banners := Get_banner()
		//其次获取阅读历史
		historys := Get_history(user_id)
		//获取热门
		hot_novels := Get_hot()
		data["banner"] = banners
		data["history"] = historys
		data["hot"] = hot_novels
		util.Return_jsonp(w, data)
	}
}

//获取banner数据的方法
func Get_banner() []models.Banner_novel {
	banner_id_list := util.Get_banner_novel_id()
	var banners []models.Banner_novel
	for _, novel_id := range banner_id_list {
		one_banner := Get_banner_by_id(novel_id)
		banners = append(banners, one_banner)
	}
	return banners
}

func Get_banner_by_id(novel_id int) models.Banner_novel {
	sql_str := "select id,name,author from book where id=" + strconv.Itoa(novel_id) + ";"
	rows, err := util.DB.Query(sql_str)
	util.CheckErr(err)
	var one_banner models.Banner_novel
	for rows.Next() {
		rows.Scan(&one_banner.Book_id, &one_banner.Name, &one_banner.Author)
	}
	return one_banner
}

//获取history的方法
func Get_history(user_id string) []models.View_history {
	//根据user_id查询view_history表,倒序排序,取前三个
	history_sql := "select book_id,chapter_id from view_history where user_id=" + user_id + " order by id desc limit 3;"
	rows, err := util.DB.Query(history_sql)
	var historys []models.View_history
	util.CheckErr(err)
	for rows.Next() {
		var one_history models.View_history
		rows.Scan(&one_history.Book_id, &one_history.Chapter_id)
		Get_book_img_name_by_id(&one_history)
		historys = append(historys, one_history)
	}
	return historys
}

func Get_book_img_name_by_id(one_history *models.View_history) {
	book_id := one_history.Book_id
	sql_str := "select name,book_img from book where id=" + strconv.Itoa(book_id) + ";"
	rows, err := util.DB.Query(sql_str)
	util.CheckErr(err)
	for rows.Next() {
		err := rows.Scan(&one_history.Name, &one_history.Image)
		util.CheckErr(err)
	}
}

//获取hot的方法
func Get_hot() []models.Hot_novel {
	var hot_novels []models.Hot_novel
	sql_str := "select id,book_img,name from book where has_chapter=1 order by rand() limit 10;"
	rows, err := util.DB.Query(sql_str)
	util.CheckErr(err)
	for rows.Next() {
		var one_hot_novel models.Hot_novel
		err := rows.Scan(&one_hot_novel.Book_id, &one_hot_novel.Image, &one_hot_novel.Name)
		util.CheckErr(err)
		Get_desc_by_book_id(&one_hot_novel)
		hot_novels = append(hot_novels, one_hot_novel)
	}
	return hot_novels
}

func Get_desc_by_book_id(one_hot_novel *models.Hot_novel) {
	book_id := one_hot_novel.Book_id
	sql_str := "select chapter_text from chapter where book_id=" + strconv.Itoa(book_id) + " limit 1;"
	rows, err := util.DB.Query(sql_str)
	util.CheckErr(err)
	var text string
	for rows.Next() {
		rows.Scan(&text)
	}
	var text_list []string
	util.Json_to_object(text, &text_list)
	var desc string
	var i = 0
	for _, sentence := range text_list {
		desc = desc + sentence
		i = i + 1
		if i > 2 {
			break
		}
	}
	one_hot_novel.Desc = desc
}
