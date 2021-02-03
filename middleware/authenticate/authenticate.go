package authenticate

import (
	"baseservice/common/session"
	"encoding/json"
	"jarvis/base/network"
	"log"
)

type (
	// 含认证中间件类型的请求
	Request struct {
		Token     string `json:"token"`      // 账号唯一标识
		Session   string `json:"session"`    // 会话标识
		SecretKey string `json:"secret_key"` // 加密 key
	}

	// 含认证中间件类型的响应
	Response struct {
		Session string `json:"session"` // 更新会话标识
	}
)

const (
	// 上下文传递额外信息键
	ContextExtraSessionKey = "Session"
)

var ()

// 校验 Session 中间件函数
func Authenticate(ctx network.Context) {
	// 反序列化数据
	request := Request{}
	if err := json.Unmarshal(ctx.Request().Data, &request); err != nil {
		if err := ctx.ServerError(err); err != nil {
			log.Printf("ctx.ServerError() error : %s", err.Error())
			return
		}
		ctx.Done()
		return
	}

	// 校验 Session
	newSession, err := session.VerifySessionAndUpdate(request.Token, request.Session, request.SecretKey)
	if err != nil {
		if err := ctx.ServerError(err); err != nil {
			log.Printf("ctx.ServerError() error : %s", err.Error())
			return
		}
		ctx.Done()
		return
	}

	// 存入上下文
	ctx.SetExtra(ContextExtraSessionKey, newSession)
}
