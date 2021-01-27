package platform

import (
	redisGo "github.com/gomodule/redigo/redis"
	"jarvis/base/database"
	"log"
	"testing"
	"time"
)

func TestQueryPlatformInfoByID(t *testing.T) {
	// 初始化 Redis
	database.InitializeRedis(time.Minute*time.Duration(5), 10, 5000, "localhost", 8000, "frank123")

	//p, err := QueryPlatformInfoByID("1")
	//if err != nil {
	//	log.Printf("%s", err.Error())
	//	return
	//}
	//
	//log.Printf("%+v", p)

	//platforms, err := QueryAllPlatformInfo()
	//if err != nil {
	//	log.Printf("%s", err.Error())
	//	return
	//}
	//
	//log.Printf("%+v", platforms)

	//i := 0
	//for i < 5000 {
	//	i++
	//	go func() {
	//		log.Println(HExistsPlatformByID("2"))
	//	}()
	//}
	//
	//time.Sleep(time.Duration(3) * time.Second)

	conn, err := database.GetRedisConn()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer conn.Close()

	log.Println(redisGo.String(conn.Do("set", "name", "frank")))
}
