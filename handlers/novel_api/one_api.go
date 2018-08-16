package novel_api

import (
	"net/http"
	"feidu/util"
	"feidu/models"
	"fmt"
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
		util.Return_json(w, data)
	}
}

func Get_chapter_name_by_book_id(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		book_id := r.FormValue("book_id")
		n := r.FormValue("n")
		sql_str := "select name from chapter where book_id =" + book_id + " limit "

		fmt.Println(n)
		fmt.Println(sql_str)
	}
}
