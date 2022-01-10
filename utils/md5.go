package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Encrypt MD5加密
func MD5Encrypt(rawString string) (result string) {

	h := md5.New()
	h.Write([]byte(rawString))
	result = hex.EncodeToString(h.Sum(nil))
	return result

}
