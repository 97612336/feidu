package util

import (
	"database/sql"
	"io/ioutil"
	"encoding/json"
	"log"
	"os/user"
	"fmt"
	"weixin_api/models"
)

var DB *sql.DB

//获取db对象
func Get_sql_db() *sql.DB {
	sqlconf := Get_conf_info()
	//打开数据库
	db, err := sql.Open("mysql",
		sqlconf.SqlUser+":"+sqlconf.SqlPassword+
			"@tcp("+sqlconf.SqlHost+":"+sqlconf.SqlPort+")/cyx?charset=utf8")
	if err != nil {
		log.Println("打开数据库出错")
	}
	return db
}

//得到家目录的路径
func Get_home_path() string {
	current_user, err := user.Current()
	if nil != err {
		fmt.Println("get user current dir err:", current_user.HomeDir)
		return ""
	}
	user_home := current_user.HomeDir
	return user_home
}

//获取mysql配置文件信息
func Get_conf_info() models.SqlConf {
	user_home := Get_home_path()
	config_file := user_home + "/conf/sqlconf"
	//	读取数据库配置文件
	data, _ := ioutil.ReadFile(config_file)
	//转化为字符串格式
	str_data := string(data)
	//实例化数据库配置类型
	var sqlconf models.SqlConf
	//得到json字符串数据
	var sql_json = []byte(str_data)
	//把json数据赋值给实例化的数据配置对象
	json.Unmarshal(sql_json, &sqlconf)
	return sqlconf
}

//获取上传文件的验证文件
func Get_img_account() models.Upload_account {
	user_home := Get_home_path()
	config_file := user_home + "/conf/upload_account"
	data, _ := ioutil.ReadFile(config_file)
	//将读取到的文件转化为字符串
	str_data := string(data)
	var account models.Upload_account
	//将读取到的配置文件赋值给类型
	var account_json = []byte(str_data)
	json.Unmarshal(account_json, &account)
	return account
}
