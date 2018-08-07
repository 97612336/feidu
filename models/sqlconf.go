package models

//定义数据库配置文件类型
type SqlConf struct {
	SqlUser     string
	SqlPassword string
	SqlHost     string
	SqlPort     string
}

//定义验证用的账号
type Upload_account struct {
	Upload_name     string
	Upload_password string
}
