package basic

import "jarvis/util/encrypt"

type ()

const (
	encryptKey = "FRANK"
)

var ()

func EncryptSecretKey(token, session string) string {
	return encrypt.MD5(token, encryptKey, session)
}
