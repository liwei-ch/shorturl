package cache

import (
	"testing"
)

func TestAddRecord(t *testing.T) {
	//AddRecord("https://www.baidu.com", "baidu")
	//res, ok := GetRecord("baidu")
	//if ok {
	//	fmt.Println(res)
	//} else {
	//	t.Fail()
	//}
	// 测试不能这么写了
	var url, surl = "http://www.baidu.com", "CAjhIbVw"
	var cache = NewCache("", "", "", "", 0)
	cache.AddUrl(url, surl)
}

func TestGetRecord(t *testing.T) {
	//AddRecord("https://www.baidu.com", "baidu")
	//res, ok := GetRecord("baidu")
	//if ok {
	//	fmt.Println(res)
	//} else {
	//	t.Fail()
	//}
	var cache = NewCache("", "", "", "", 0)
	var url, surl = "http://www.baidu.com", "CAjhIbVw"
	cache.AddUrl(url, surl)
	for i := 0; i < 1000; i++ {
		res, ok := cache.GetUrl(surl)

		if !ok || res != url {
			t.Fail()
		}
	}
}
