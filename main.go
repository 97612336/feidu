package main

import (
	"net/http"
	"log"
	"weixin_api/handlers"
	_ "github.com/go-sql-driver/mysql"
	)

func init(){
	//util.DB=util.Get_sql_db()
}



func main(){
	//defer util.DB.Close()

	var err error
	//设置路由
	handlers.MyUrls()
	//设置监听端口
	err = http.ListenAndServe(":8080",nil)
	//启动程序
	if err !=nil{
		log.Fatal("出现错误：",err)
	}
	
}
