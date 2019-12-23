package models

import (
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3" // 具体的数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//设置 gofmt需要等待一下，不会立即生效

//const (
//	_DB_NAME        = "data/beelog.db"
//	_SQLITE3_DRIVER = "sqlite3"
//)

//func RegisterDB() {
//	if !com.IsExist(_DB_NAME) {
//		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
//		os.Create(_DB_NAME)
//	}
//	orm.RegisterModel(new(Category), new(Topic))
//	orm.RegisterDriver(_DB_NAME, orm.DRSqlite)
//	// 必须要有一个叫default的数据库
//	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10) // 10是连接池大小
//}

type Category struct {
	Id          int64
	Title       string
	CreatedTime time.Time `orm:"index;auto_now_add"`
	Views       int64     `orm:"index"`
	//auto_now_add相当于mysql原生语句： date default(CURRENT_TIMESTAMP)
	TopicTime       time.Time `orm:"index;auto_now_add"`
	TopicCount      int64     // Topic是文章的意思
	TopicLastUserId int64
}

type Topic struct {
	Id         int64
	Title      string
	Uid        int64
	Views      int64
	Content    string `orm:"size(5000)"`
	Attachment string

	CreatedTime     time.Time `orm:"index"`
	UpdatedTime     time.Time `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64     `orm:"index"`
	ReplyLastUserId int64     `orm:"index"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Category), new(Topic))
}
