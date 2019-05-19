package url

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"shorturl/url/cache"
	"strconv"
	"strings"
	"time"
)

const (
	_CHARS    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_CHAR_LEN = len(_CHARS)
)

type pair struct {
	url  string
	surl string
}

type UrlServer struct {
	c      *cache.Cache
	prefix string
	ch     chan pair
}

func NewUrlServer(redisAddr, redisPass, redisKey, sqlServer, prefix string, redisDbNum int) *UrlServer {
	var server UrlServer
	server.c = cache.NewCache(redisAddr, redisPass, redisKey, sqlServer, redisDbNum)
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	server.prefix = prefix
	server.ch = make(chan pair, 100)

	for i := 0; i < 50; i++ {
		// 启动goroutine
		go server.addCache()
	}

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

func (u *UrlServer) addCache() {
	for {
		p, ok := <-u.ch
		if !ok {
			return
		}
		u.c.AddUrl(p.url, p.surl)
	}
}

func (u *UrlServer) GenerateUrl(url string) string {
	surl := generate(url)
	//fmt.Println(surl)
	// 放到goroutine中做
	// goroutine会导致效率下降
	//go u.c.AddUrl(url, surl)
	u.ch <- pair{url, surl}
	return u.prefix + surl
}

func (u *UrlServer) Close() error {
	for {
		select {
		case p, ok := <-u.ch:
			if !ok {
				break
			}
			u.c.AddUrl(p.surl, p.surl)
		case <-time.After(10 * time.Second):
			close(u.ch)
			return nil
		}
	}
}
