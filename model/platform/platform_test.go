package platform

import (
	"jarvis/base/database"
	"log"
	"testing"
	"time"
)

func TestQueryPlatformInfoByID(t *testing.T) {
	// 初始化 Redis
	database.InitializeRedis(time.Minute*time.Duration(5), 10, 30, "localhost", 6379, "frank123")

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

	log.Println(HExistsPlatformByID("3"))
}
