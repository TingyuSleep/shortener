package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Sum 对传入的参数求md5值
// 通过加密散列函数，产生出一个128位（16字节）的散列值（通常表现为一个32位的十六进制数字）
func Sum(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil)) // 32位十六进制数
}
