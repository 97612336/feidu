package novel_api

import (
	"net/http"
	"feidu/util"
	"feidu/models"
	"strconv"
)

func Show_all_categories(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})
		//得到所有小说的分类
		categorys := Get_all_novel_category()

		data["categorys"] = categorys
		data["code"] = 200
		util.Return_jsonp(w, data)

	}
}

//得到所有小说的分类
func Get_all_novel_category() []models.Category {
	sql_str := "SELECT * from category;"
	rows, err := util.DB.Query(sql_str)
	defer rows.Close()
	util.CheckErr(err)
	var categorys []models.Category
	for rows.Next() {
		var one_category models.Category
		rows.Scan(&one_category.CategoryId, &one_category.CategoryName, &one_category.CategoryImage)
		categorys = append(categorys, one_category)
	}
	return categorys
}

func Show_category_book(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		var data = make(map[string]interface{})
		//接受参数
		n := util.Get_argument(r, "n")
		category_id := util.Get_argument(r, "category_id")
		//根据页数和种类id查询得到书籍结果
		books := Get_category_book(n, category_id)
		data["code"] = 200
		data["data"] = books
		util.Return_jsonp(w, data)
	}
}

func Get_category_book(n string, category_id string) []models.Category_book {
	var category_books []models.Category_book
	page_size := 20
	sql_str := "select id,book_img,name from book where kind_id=? and has_chapter=1 limit ?,?"
	page, err := strconv.Atoi(n)
	util.CheckErr(err)
	start_num := (page - 1) * page_size
	rows, err := util.DB.Query(sql_str, category_id, start_num, page_size)
	defer rows.Close()
	util.CheckErr(err)
	for rows.Next() {
		var one_category_book models.Category_book
		rows.Scan(&one_category_book.BookId, &one_category_book.Image, &one_category_book.Name)
		desc := Get_book_desc_by_book_id(one_category_book.BookId)
		one_category_book.Desc = desc
		category_books = append(category_books, one_category_book)
	}
	return category_books
}

//根据book_id获取描述
func Get_book_desc_by_book_id(book_id int) string {
	//查询chapter表得到描述
	sql_str := "select chapter_text from chapter where book_id=? limit 1;"
	rows, err := util.DB.Query(sql_str, book_id)
	defer rows.Close()
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
	return desc
}
