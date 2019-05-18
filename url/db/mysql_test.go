package db

import (
	"testing"
)

// 这些函数都是有状态的，不太好测试
func TestAddRecord(t *testing.T) {
	//t.Logf("%s\n", strconv.Quote(""))
	//err := AddRecord("' OR TRUE", "")
	//if err != INVALID_PARAM {
	//	t.Fail()
	//}
	//err = AddRecord("", "")
	//if err != INVALID_PARAM {
	//	t.Fail()
	//}
	//
	//err = AddRecord("https://baidu.com", "baidu")
	//if err != nil && err != RECORD_EXISTED {
	//	t.Fail()
	//}
	var url, surl = "http://www.baidu.com", "CAjhIbVw"

	var db = NewMysqlDB("")
	err := db.AddRecord(Record{Surl: surl, Url: url})
	if err != nil {
		t.Fail()
	}
	res, err := db.GetRecord(surl)
	if err != nil || res.Url != url {
		t.Fail()
	}
}

func TestGetRecord(t *testing.T) {
	//res, err := GetRecord("null")
	//if err != NO_RECORD {
	//	t.Fail()
	//}
	//res, err = GetRecord("' OR TRUE")
	//if err != INVALID_PARAM {
	//	t.Fail()
	//}
	//
	//res, err = GetRecord("baidu")
	//if err != nil || res != "https://baidu.com" {
	//	t.Fail()
	//}
	var url, surl = "http://www.baidu.com", "CAjhIbVw"
	var db = NewMysqlDB("")
	_ = db.AddRecord(Record{Surl: surl, Url: url})

	for i := 0; i < 100; i++ {
		res, err := db.GetRecord(surl)
		if err != nil || res.Url != url {
			t.Fail()
		}
	}
}
