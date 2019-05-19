package cache

import (
	"net/url"
	"shorturl/url/db"

	"github.com/go-redis/redis"
)

type Cache struct {
	clnt *redis.Client
	key  string
	db   *db.MysqlDB
}

func NewCache(addr, pass, key, sqlServer string, dbNum int) *Cache {
	var c Cache
	c.clnt = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       dbNum,
	})
	c.key = key
	c.db = db.NewMysqlDB(sqlServer)

	return &c
}

func (c *Cache) AddUrl(originUrl, surl string) {
	var record db.Record
	record.Surl = surl
	record.Url = url.QueryEscape(originUrl)
	// 先检查有没有
	res, err := c.clnt.HGet(c.key, surl).Result()
	// EOF错误是什么？
	// 已解决：redis连接被服务器断开了，将服务器redis连接时间设置一下就好了
	if err == nil {
		// 已被缓存并且缓存的和参数相等
		if res == record.Url {
			return
		}
		// 已缓存的和url不等
		_ = c.db.AddRecord(record)
	} else if err == redis.Nil {
		c.clnt.HSet(c.key, record.Surl, record.Url)
		// 数据库中可能有，可能没有，直接加入或者更新
		_ = c.db.AddRecord(record)
	} else {
		// 其它未知错误
		panic(err)
	}
}

// 查找这个没办法goroutine省时间
func (c *Cache) GetUrl(surl string) (string, bool) {
	res, err := c.clnt.HGet(c.key, surl).Result()
	if err == nil {
		res, _ = url.QueryUnescape(res)
		return res, true
	} else if err == redis.Nil {
		// 查找数据库
		record, err := c.db.GetRecord(surl)
		if err == nil {
			// 写入缓存
			//go c.AddUrl(record.Url, surl)
			c.clnt.HSet(c.key, record.Surl, record.Url)
			res, _ := url.QueryUnescape(record.Url)
			return res, true
		}
	}
	// 都没查找到
	return "nil", false
}
