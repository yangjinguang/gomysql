package gomysql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func Conn(url string) (db *DB, err error) {
	db = &DB{}
	sqlDB, err := sql.Open("mysql", url)
	if err != nil {
		return db, err
	}
	db.DB = sqlDB
	return db, nil
}
