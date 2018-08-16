package main

import (
	"net/http"
	"feidu/handlers"
	_ "github.com/go-sql-driver/mysql"
	"feidu/util"
)

func init() {
	util.DB = util.Get_sql_db()
}

func main() {
	defer util.DB.Close()
	//设置路由
	handlers.MyUrls()
	//设置监听端口
	err := http.ListenAndServe(":8080", nil)
	//启动程序
	util.CheckErr(err)

}
