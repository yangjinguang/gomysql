package gomysql

import (
	"testing"
	"fmt"
	"time"
)

var db *DB

type TmpUser struct {
	Id        int64     `json:"id" mysql:"id"`
	NickName  string    `json:"nick_name" mysql:"nick_name"`
	CreatedAt time.Time `json:"created_at" mysql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" mysql:"updated_at"`
}

func init() {
	d, err := Conn(fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?parseTime=true",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"test_db",
	))
	db = d
	if err != nil {
		panic(err.Error())
	}
}

func TestDB_All(t *testing.T) {
	var users []*TmpUser

	err := db.T("users").
		Select().
		Where("").
		Limit(0, 1).
		All(&users)
	if err != nil {
		t.Fail()
	}
}

func TestDB_One(t *testing.T) {
	user := TmpUser{}
	_, err := db.T("users").
		SelectById(88).
		One(&user)
	if err != nil {
		t.Fail()
	}
}

func TestDB_Insert(t *testing.T) {
	user := TmpUser{}
	user.NickName = "test-444"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := db.T("users").Insert(&user)
	if err != nil {
		t.Fail()
	}
}

func TestDB_Replace(t *testing.T) {
	user := TmpUser{}
	user.Id = 1
	user.NickName = "test-444"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	db.T("users").Replace(&user)
}

func TestDB_Delete(t *testing.T) {
	err := db.T("users").Where("`nick_name` = 'test-444'").Delete()
	if err != nil {
		t.Fail()
	}
}
