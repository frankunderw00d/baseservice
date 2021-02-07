package traceRecord

import (
	"baseservice/middleware/authenticate"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"jarvis/base/database"
	"jarvis/base/network"
	"log"
	"time"
)

type (
	// 追踪请求记录类型
	RecordType int
)

const (
	// 追踪请求记录表名
	CollectionName = "trace_record"

	// 追踪普通请求记录类型
	RecordTypeNormal = 0
	// 追踪需要校验记录类型
	RecordTypeAuthenticate = 1

	// 追踪请求记录类型名
	RecordTypeField = "type"
	// 追踪普通请求记录字段名
	RecordModuleField = "module"
	RecordRouteField  = "route"
	RecordReplyField  = "reply"
	RecordIDField     = "id"
	RecordDataField   = "data"
	// 追踪需要校验的请求记录字段名
	RecordTokenField     = "token"
	RecordSessionField   = "session"
	RecordSecretKeyField = "secretKey"
)

var ()

func init() {

}

// 普通追踪请求
func TraceRecord(ctx network.Context) {
	request := ctx.Request()

	if err := record(
		request.Module,
		request.Route,
		request.Reply,
		request.ID,
		string(request.Data),
		RecordTypeNormal,
		"",
		"",
		"",
	); err != nil {
		if err := ctx.ServerError(err); err != nil {
			log.Printf("ctx.ServerError() error : %s", err.Error())
			return
		}
		ctx.Done()
	}
}

// 追踪需要校验的请求
func TraceAuthenticateRecord(ctx network.Context) {
	// 反序列化数据
	request := authenticate.Request{}
	if err := json.Unmarshal(ctx.Request().Data, &request); err != nil {
		if err := ctx.ServerError(err); err != nil {
			log.Printf("ctx.ServerError() error : %s", err.Error())
			return
		}
		ctx.Done()
		return
	}

	if err := record(
		ctx.Request().Module,
		ctx.Request().Route,
		ctx.Request().Reply,
		ctx.Request().ID,
		string(ctx.Request().Data),
		RecordTypeAuthenticate,
		request.Token,
		request.Session,
		request.SecretKey,
	); err != nil {
		if err := ctx.ServerError(err); err != nil {
			log.Printf("ctx.ServerError() error : %s", err.Error())
			return
		}
		ctx.Done()
	}
}

// 记录
func record(module, route, reply, id, data string, recordType RecordType, token, session, secretKey string) error {
	conn, err := database.GetMongoConn(CollectionName)
	if err != nil {
		return err
	}

	_, err = conn.InsertOne(context.Background(), bson.M{
		RecordModuleField:    module,
		RecordRouteField:     route,
		RecordReplyField:     reply,
		RecordIDField:        id,
		RecordDataField:      data,
		RecordTypeField:      recordType,
		RecordTokenField:     token,
		RecordSessionField:   session,
		RecordSecretKeyField: secretKey,
		"time":               time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return err
	}
	return nil
}
