package basic

import "jarvis/util/encrypt"

type ()

const (
	// 加密密钥
	encryptKey = "FRANK"
)

var ()

// 对用户 Token 和 Session 加上加密密钥进行加密
func EncryptSecretKey(token, session string) string {
	return encrypt.MD5(token, encryptKey, session)
}
