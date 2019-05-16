package db

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var (
	db  *sql.DB
	err error
)

var (
	NO_RECORD      = errors.New("no record found")
	INVALID_PARAM  = errors.New("invalid parameter")
	RECORD_EXISTED = errors.New("record existed")
	//RECORD_UPDATED = errors.New("record updated")
)

func init() {
	//fmt.Println("mysql inited")
	//io.EOF
	db, err = sql.Open("mysql", "<put your mysql url>")
	// 连接何时关闭？
	// defer db.Close()
	if err != nil {
		// 连接不上直接崩掉
		panic(err)
	}
}

func AddRecord(url, surl string) error {
	// 如何保证不被sql注入？？？
	if surl == "" || strings.ContainsAny(surl, "' ") {
		// 拒绝插入
		return INVALID_PARAM
	}
	rows, err := db.Query("SELECT url FROM records WHERE surl = ?", surl)
	if err != nil {
		return nil
	}

	if rows.Next() {
		//fmt.Println("have row in database")
		var db_url string
		err = rows.Scan(&db_url)
		if err != nil {
			return nil
		}
		//fmt.Println("db url: ", db_url)
		if db_url != url {
			// 更新链接
			//fmt.Println("update for ", surl)
			_, err := db.Exec("UPDATE records SET url = ? WHERE surl = ?", url, surl)
			if err != nil {
				return err
			}
			//fmt.Println(res.RowsAffected())
			// 这个感觉不能算错误吧
			return nil
		}
		// 在此停止
		return RECORD_EXISTED
	}

	_, err = db.Exec("INSERT INTO records(surl, url) VALUE (?, ?)", surl, url)
	return err
}

func GetRecord(surl string) (string, error) {
	// 检测参数有效性
	if surl == "" || strings.ContainsAny(surl, "' ") {
		return "nil", INVALID_PARAM
	}
	// 查询链接
	rows, err := db.Query("SELECT url FROM records WHERE surl = ?", surl)

	if err != nil {
		return "nil", err
	}

	if rows.Next() {
		var res string
		err = rows.Scan(&res)
		if err != nil {
			panic(err)
		}

		return res, nil
	}
	// 查找为空
	return "nil", NO_RECORD
}
