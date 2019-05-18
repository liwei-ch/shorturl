package url

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"shorturl/url/cache"
	"strconv"
	"strings"
)

const (
	_CHARS    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_CHAR_LEN = len(_CHARS)
)

type UrlServer struct {
	c      *cache.Cache
	prefix string
}

func NewUrlServer(redisAddr, redisPass, redisKey, sqlServer, prefix string, redisDbNum int) *UrlServer {
	var server UrlServer
	server.c = cache.NewCache(redisAddr, redisPass, redisKey, sqlServer, redisDbNum)
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	server.prefix = prefix
	return &server
}

func generate(url string) string {
	sum := md5.New()
	sum.Write([]byte(url))
	code := fmt.Sprintf("%x", sum.Sum(nil))
	var buf = bytes.NewBuffer(nil)
	// 产生链接
	for i := 0; i < 32; i += 4 {
		val, _ := strconv.ParseUint(code[i:i+4], 16, 32)
		buf.WriteByte(_CHARS[val%uint64(_CHAR_LEN)])
	}
	// 只产生映射后的字符串，不参与持久化
	return buf.String()
}

func (u *UrlServer) GenerateUrl(url string) string {
	str := generate(url)
	//fmt.Println(str)
	// 放到goroutine中做
	go u.c.AddUrl(url, str)
	return u.prefix + str
}
