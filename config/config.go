package config

import "weixin_api/util"

func Get_path() string {
	home_path := util.Get_home_path()
	return home_path
}

//静态文件存放的本地地址
var Static_Path = Get_path() + "/static_dir/"
//静态文件存放目录二
var Static_path_web = "/static_dir/"

//静态文件读取的网络地址前缀
var Static_href = "http://140.143.224.74:3456/"

//本地静态文件目录
var Local_href = "http://127.0.0.1:3456/"
