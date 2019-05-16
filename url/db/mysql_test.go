package db

import (
	"testing"
)

func TestAddRecord(t *testing.T) {
	//t.Logf("%s\n", strconv.Quote(""))
	err := AddRecord("' OR TRUE", "")
	if err != INVALID_PARAM {
		t.Fail()
	}
	err = AddRecord("", "")
	if err != INVALID_PARAM {
		t.Fail()
	}

	err = AddRecord("https://baidu.com", "baidu")
	if err != nil && err != RECORD_EXISTED {
		t.Fail()
	}
}

func TestGetRecord(t *testing.T) {
	res, err := GetRecord("null")
	if err != NO_RECORD {
		t.Fail()
	}
	res, err = GetRecord("' OR TRUE")
	if err != INVALID_PARAM {
		t.Fail()
	}

	res, err = GetRecord("baidu")
	if err != nil || res != "https://baidu.com" {
		t.Fail()
	}
}
