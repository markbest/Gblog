package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

//字符串类型转化为int类型
func StringToInt(str string) (id int) {
	id, _ = strconv.Atoi(str)
	return id
}

//字符串类型转化为int64类型
func StringToInt64(str string) (id int64) {
	id, _ = strconv.ParseInt(str, 10, 64)
	return id
}

//int64类型转化为字符串类型
func Int64ToString(num int64) (str string) {
	str = strconv.FormatInt(num, 10)
	return str
}

//md5方法
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//字串截取
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//检查是否是当前URL
func IsActive(request_url string, target_url string) bool {
	var is_active bool
	url_str, _ := url.Parse(request_url)

	// 分类active
	if url_str.Path == "/category/"+target_url {
		is_active = true
		return is_active
	} else {
		is_active = false
	}

	//后台导航栏active
	if url_str.Path == target_url {
		is_active = true
		return is_active
	} else {
		is_active = false
	}

	//导航栏下一级页面active
	if strings.Contains(url_str.Path, target_url) {
		is_active = true
		return is_active
	} else {
		is_active = false
	}
	return is_active
}

//获取静态文件的版本号
func GetStaticVersion(filename string) (t string) {
	file_path := "static/css/" + filename
	if fileInfo, err := os.Stat(file_path); err == nil {
		t = fileInfo.ModTime().Format("200601021504")
	} else {
		fmt.Println(err)
	}
	return t
}

//获取
func GetBaseUrl(prefix string, url string) string {
	return beego.AppConfig.String("base_url") + prefix + url
}
