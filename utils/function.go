package utils

import (
	"encoding/hex"
	"crypto/md5"
)

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

//检查是否是当前URL
func IsActive(request_url string, target_url string) bool {
	var is_active bool
	if request_url == "/category/" + target_url {
		is_active = true
		return is_active
	} else {
		is_active =  false
	}

	if request_url == target_url {
		is_active = true
		return is_active
	} else {
		is_active =  false
	}
	return is_active
}