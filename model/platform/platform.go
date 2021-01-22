// 平台信息以 hash map 的方式存在 Redis ，
package platform

import (
	"baseservice/base/basic"
	"encoding/json"
	"errors"
	"fmt"
	redisGo "github.com/gomodule/redigo/redis"
	"jarvis/base/database"
	"time"
)

type (
	// 平台表映射结构
	Platform struct {
		ID       int64     `json:"id"`        // id
		Name     string    `json:"name"`      // 平台名
		Link     string    `json:"link"`      // 平台连接
		Owner    string    `json:"owner"`     // 平台拥有者
		CreateAt time.Time `json:"create_at"` // 创建时间
		UpdateAt time.Time `json:"update_at"` // 修改时间
	}

	// 平台列表
	PlatformList []Platform
)

const (
	// 平台表对应的表名
	MySQLPlatformTableName = "static_platform"
	// 平台信息哈希字典名
	DefaultPlatformKey basic.ComposeString = "GLOBAL:PLATFORM"
	// 平台信息前缀
	DefaultPlatformFieldPrefix basic.ComposeString = "id:"
)

var ()

func (p *Platform) MySQLTableName() string {
	return MySQLPlatformTableName
}

func (pl PlatformList) QueryOrder() string {
	return fmt.Sprintf("select * from %s", MySQLPlatformTableName)
}

// 根据 id 查询平台信息
func QueryPlatformInfoByID(id string) (Platform, error) {
	conn, err := database.GetRedisConn()
	if err != nil {
		return Platform{}, err
	}
	defer conn.Close()

	str, err := redisGo.String(conn.Do("hget", DefaultPlatformKey.String(), DefaultPlatformFieldPrefix.Compose(id)))
	if err != nil {
		return Platform{}, err
	}

	p := Platform{}
	if err := json.Unmarshal([]byte(str), &p); err != nil {
		return Platform{}, err
	}

	return p, nil
}

// 查询所有平台信息
func QueryAllPlatformInfo() ([]Platform, error) {
	conn, err := database.GetRedisConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reply, err := conn.Do("hgetall", DefaultPlatformKey.String())
	if err != nil {
		return nil, err
	}

	bytesList, ok := reply.([]interface{})
	if !ok {
		return nil, errors.New("return type error")
	}

	platforms := make([]Platform, 0)
	for index, data := range bytesList {
		if index%2 == 0 {
			continue
		}

		platform := Platform{}
		if err := json.Unmarshal(data.([]byte), &platform); err != nil {
			return nil, err
		}
		platforms = append(platforms, platform)
	}

	return platforms, nil
}

// 设置平台信息到 Redis 的哈希表中，key 为 DefaultPlatformKey
func HSetPlatformInfoByID(id string, value string) error {
	conn, err := database.GetRedisConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = redisGo.Int(conn.Do("hset", DefaultPlatformKey.String(), DefaultPlatformFieldPrefix.Compose(id), value))
	if err != nil {
		return err
	}

	return nil
}

// 验证平台id是否存在
func HExistsPlatformByID(id string) bool {
	conn, err := database.GetRedisConn()
	if err != nil {
		return false
	}
	defer conn.Close()

	reply, err := redisGo.Int(conn.Do("hexists", DefaultPlatformKey.String(), DefaultPlatformFieldPrefix.Compose(id)))
	if err != nil {
		return false
	}

	return reply == 1
}
