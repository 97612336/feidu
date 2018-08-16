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
		util.Return_json(w, data)
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
		util.Return_json(w, data)
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
		util.Return_json(w, data)
	}
}
