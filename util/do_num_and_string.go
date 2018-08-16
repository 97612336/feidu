package util

import (
	"time"
	"math/rand"
	"crypto/md5"
	"io"
	"encoding/hex"
)

//获取一组随机数
func Get_random_arr(count int, max_num int) []int {
	//定义随机数返回数组
	var nums_arr []int
	//定义一个中间值int64数字
	var one_tmp_num int64 = 1
	//如果数组大小小于规定的长度,则执行遍历
	for len(nums_arr) < count {
		//定义随机数种子
		time_int := time.Now().Unix() + one_tmp_num
		rand.Seed(time_int)
		//生成随机数
		one_rand_num := rand.Intn(max_num)
		one_tmp_num = time_int
		nums_arr = append(nums_arr, one_rand_num)
	}
	return nums_arr
}


//得到md5字符串
func Get_md5str(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
