package url

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"shorturl/url/cache"
	"strconv"
)

const (
	_CHARS    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_CHAR_LEN = len(_CHARS)
	_PREFIX   = "http://rammiah.org:8080/to/"
)

func generate(url string) (string) {
	sum := md5.New()
	sum.Write([]byte(url))
	code := fmt.Sprintf("%x", sum.Sum(nil))
	var buf = bytes.NewBuffer(nil)
	// 产生链接
	for i := 0; i < 32; i += 4 {
		val, _ := strconv.ParseUint(code[i:i+4], 16, 32)
		buf.WriteByte(_CHARS[val%uint64(_CHAR_LEN)])
	}

	// 放到goroutine中做
	go cache.AddRecord(url, buf.String())

	return buf.String()
}

func GenerateUrl(url string) string {
	str := generate(url)
	//fmt.Println(str)
	return _PREFIX + str
}
