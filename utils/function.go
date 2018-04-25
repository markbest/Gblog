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
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//检查是否是当前URL
func IsActive(requestUrl string, targetUrl string) bool {
	var isActive bool
	urlStr, _ := url.Parse(requestUrl)

	// 分类active
	if urlStr.Path == "/category/"+targetUrl {
		isActive = true
		return isActive
	} else {
		isActive = false
	}

	//后台导航栏active
	if urlStr.Path == targetUrl {
		isActive = true
		return isActive
	} else {
		isActive = false
	}

	//导航栏下一级页面active
	if strings.Contains(urlStr.Path, targetUrl) {
		isActive = true
		return isActive
	} else {
		isActive = false
	}
	return isActive
}

//获取静态文件的版本号
func GetStaticVersion(filename string) (t string) {
	filePath := "static/css/" + filename
	if fileInfo, err := os.Stat(filePath); err == nil {
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
