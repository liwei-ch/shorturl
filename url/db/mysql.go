package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDB struct {
	db *gorm.DB
}

func NewMysqlDB(sqlServer string) *MysqlDB {
	sqlDb, err := gorm.Open("mysql", sqlServer)
	if err != nil {
		panic(err)
	}
	return &MysqlDB{db: sqlDb}
}

type Record struct {
	Surl string `gorm:"column:surl;primary_key;not_null"`
	Url  string `gorm:"column:url;not_null"`
}

func (Record) TableName() string {
	return "records"
}

var (
	NoRecord      = errors.New("no record found")
	InvalidParam  = errors.New("invalid parameter")
	RecordExisted = errors.New("record existed")
)

func (m *MysqlDB) AddRecord(record Record) error {
	// 先查找是否已有
	rec, err := m.GetRecord(record.Surl)
	if err == NoRecord {
		return m.db.Save(&record).Error
	} else if err == nil {
		if rec.Url != record.Url {
			m.db.Table(rec.TableName()).Where("surl = ?", record.Surl).Update("url", record.Url)
		}
	}
	return nil
}

func (m *MysqlDB) GetRecord(surl string) (record Record, err error) {
	var dbErr = m.db.Where("surl = ?", surl).First(&record)
	if dbErr.Error != nil {
		if dbErr.Error == gorm.ErrRecordNotFound {
			err = NoRecord
			return
		}
		//panic(dbErr.Error)
		err = dbErr.Error
		return
	}
	return
}
