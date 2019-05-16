package cache

import (
	"shorturl/url/db"

	"github.com/go-redis/redis"
)

var (
	clnt *redis.Client
	key  = "records"
)

func init() {
	//fmt.Println("redis inited")
	clnt = redis.NewClient(&redis.Options{
		Addr:     "<ip>:<port>", // put redis server  address
		Password: "<password>",  // put redis password
		DB:       0,
	})
}

func AddRecord(url, surl string) {
	// 先检查有没有
	res, err := clnt.HGet(key, surl).Result()
	// EOF错误是什么？
	if err == nil {
		// 已被缓存并且缓存的和参数相等
		if res == url {
			return
		}
		// 已缓存的和url不等
		err = db.AddRecord(url, surl)
		if err != nil {
			panic(err)
		}
	} else if err == redis.Nil {
		clnt.HSet(key, surl, url)
		// 数据库中可能有
		_, err := db.GetRecord(surl)
		if err == db.NO_RECORD {
			// 插入
			err = db.AddRecord(url, surl)
			if err != nil {
				panic(err)
			}
		}
	} else {
		panic(err)
	}
}

// 查找这个没办法goroutine省时间
func GetRecord(surl string) (string, bool) {
	res, err := clnt.HGet(key, surl).Result()
	if err == nil {
		return res, true
	} else if err == redis.Nil {
		// 查找数据库
		db_url, err := db.GetRecord(surl)
		if err == nil {
			// 写入缓存
			go AddRecord(db_url, surl)
			return db_url, true
		}
	}

	return "nil", false
}
