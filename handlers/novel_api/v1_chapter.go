package novel_api

import (
	"net/http"
	"feidu/util"
	"feidu/models"
		)

func Get_one_chapter_by_book_id(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})
		book_id := util.Get_argument(r, "book_id")
		user_id := util.Get_argument(r, "user_id")
		//根据book_id和user_id查询阅读记录表，如果没有就返回当前书本的第一章
		one_chapter := Read_book(book_id, user_id)
		//根据book_id查询出书本信息
		book_name := Get_book_name(book_id)
		data["code"] = 200
		data["book_id"] = book_id
		data["book_name"] = book_name
		data["chapter"] = one_chapter
		util.Return_jsonp(w, data)
	}

}

func Read_book(book_id string, user_id string) models.One_chapter {
	sql_str := "select chapter_id from view_history where book_id=? and user_id=? order by id desc limit 1;"
	rows, err := util.DB.Query(sql_str, book_id, user_id)
	defer rows.Close()
	util.CheckErr(err)
	var chapter_id int
	for rows.Next() {
		err := rows.Scan(&chapter_id)
		util.CheckErr(err)
	}
	//如果chapter_id=0,则返回本书的第一节，否则，返回本书的chapter_id章节
	if chapter_id == 0 {
		book_sql := "select id,name,chapter_text from chapter where book_id=? limit 1;"
		rows, err := util.DB.Query(book_sql, book_id)
		defer rows.Close()
		util.CheckErr(err)
		var one_chapter models.One_chapter
		for rows.Next() {
			var text string
			err := rows.Scan(&one_chapter.ChapterId, &one_chapter.ChapterName, &text)
			util.CheckErr(err)
			var text_list []string
			util.Json_to_object(text, &text_list)
			one_chapter.ChapterContent = text_list
		}
		return one_chapter
	} else {
		book_sql := "select id,name,chapter_text from chapter where id=?;"
		rows, err := util.DB.Query(book_sql, chapter_id)
		defer rows.Close()
		util.CheckErr(err)
		var one_chapter models.One_chapter
		for rows.Next() {
			var text string
			err := rows.Scan(&one_chapter.ChapterId, &one_chapter.ChapterName, &text)
			util.CheckErr(err)
			var text_list []string
			util.Json_to_object(text, &text_list)
			one_chapter.ChapterContent = text_list
		}
		return one_chapter
	}
}

func Get_book_name(book_id string) string {
	sql_str := "select name from book where id=?;"
	rows, err := util.DB.Query(sql_str, book_id)
	defer rows.Close()
	util.CheckErr(err)
	var book_name string
	for rows.Next() {
		err := rows.Scan(&book_name)
		util.CheckErr(err)
	}
	return book_name
}
