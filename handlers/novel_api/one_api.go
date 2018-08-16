package novel_api

import (
	"net/http"
	"feidu/util"
	"feidu/models"
	"fmt"
	"strconv"
)

func Get_some_book_name(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})
		var books []models.Book
		//得到一组随机数
		ids_arr := util.Get_random_arr(20, 17339)
		fmt.Println(ids_arr)
		for _, one_id := range ids_arr {
			fmt.Println(one_id)
			new_book := Get_one_book_by_id(one_id)
			books = append(books, new_book)
		}
		data["books"] = books
		util.Return_json(w, data)
	}
}

//根据一本书的id获取整本数的信息
func Get_one_book_by_id(one_id int) models.Book {
	sql_str := "select * from book where id >" + strconv.Itoa(one_id) + " limit 1;"
	fmt.Println(sql_str)
	rows, err := util.DB.Query(sql_str)
	util.CheckErr(err)
	var one_book models.Book
	for rows.Next() {
		err := rows.Scan(&one_book.Id, &one_book.Name, &one_book.Create_time, &one_book.Url,
			&one_book.Book_img, &one_book.Kind, &one_book.Author, &one_book.Has_chapter)
		util.CheckErr(err)
	}
	return one_book
}
