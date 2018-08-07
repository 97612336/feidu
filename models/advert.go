package models

import "time"

type Advert struct {
	Name string	//广告名称
	Title string	//广告标题
	Img string	//广告图片
	Media string	//广告媒体
	Link string	//广告的推广链接
	OnlineTime time.Time	//广告的上线时间
	OffLineTime time.Time	//广告的下线时间
//	广告类型，1表示横幅通栏上
//	2表示横幅通栏下
//	3表示服务
//	4表示插屏
//	5表示发现
//	6表示开屏
//	7表示热销车
	Kind int 		//广告的类型
}
