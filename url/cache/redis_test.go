package cache

import (
	"fmt"
	"testing"
)

func TestAddRecord(t *testing.T) {
	AddRecord("https://www.baidu.com", "baidu")
	res, ok := GetRecord("baidu")
	if ok {
		fmt.Println(res)
	} else {
		t.Fail()
	}
}

func TestGetRecord(t *testing.T) {
	AddRecord("https://www.baidu.com", "baidu")
	res, ok := GetRecord("baidu")
	if ok {
		fmt.Println(res)
	} else {
		t.Fail()
	}
}
