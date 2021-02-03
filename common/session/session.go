package session

import (
	"baseservice/base/basic"
	"errors"
	"jarvis/base/database/redis"
	uRand "jarvis/util/rand"
	"strings"
	"time"
)

const (
	// 用户 Session 键
	UsersSessionKey = "UsersSession"
	// 用户 Session 列键
	UsersSessionField basic.ComposeString = "User:"
)

// 设置 Session
func SetSession(token, session string) error {
	now := time.Now().Format("20060102150405")
	_, err := redis.HSet(UsersSessionKey, UsersSessionField.Compose(token), session+":"+now)
	if err != nil {
		return err
	}

	return nil
}

// 获取 Session
func GetSession(token string) (string, error) {
	v, err := redis.HGet(UsersSessionKey, UsersSessionField.Compose(token))
	if err != nil {
		return "", err
	}

	return strings.SplitN(v, ":", 2)[0], nil
}

// Session 是否超时，默认15分钟
func CheckSessionTimeout(token string) (bool, error) {
	v, err := redis.HGet(UsersSessionKey, UsersSessionField.Compose(token))
	if err != nil {
		return false, err
	}

	p, err := time.ParseInLocation("20060102150405", strings.SplitN(v, ":", 2)[1], time.Local)
	if err != nil {
		return false, err
	}

	if time.Now().Sub(p).Minutes() >= 15 {
		return true, nil
	}

	return false, nil
}

// 验证 Session 并更新
// 如果未过期则返回 ""
// 过期则返回更新后的 Session
func VerifySessionAndUpdate(token, session, secretKey string) (string, error) {
	// 根据 request.Token 取得 Session
	redisSession, err := GetSession(token)
	if err != nil {
		return "", err
	}

	// 核对 Session 和 secretKey
	if redisSession != session {
		return "", errors.New("session wrong")
	}
	if basic.EncryptSecretKey(token, redisSession) != secretKey {
		return "", errors.New("secretKey wrong")
	}

	// 核对 Session 是否超时，超时则更换，存入 Redis 且返回给用户
	timeout, err := CheckSessionTimeout(token)
	if err != nil {
		return "", err
	}

	newSession := session

	// 超时
	if timeout {
		// 生成新的 Session
		newSession = uRand.RandomString(8)
		// 保存新的 Session 到 redis 中
		err := SetSession(token, newSession)
		if err != nil {
			return "", err
		}
	} else {
		// 未超时返回 ""
		newSession = ""
	}

	return newSession, nil
}