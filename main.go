package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"

	//_ "myBeego/models"
	_ "myBeego/routers"
)

//func init() {
//	models.RegisterDB()
//}

//var BlodDb orm.Ormer

type Topic345 struct {
	Id    int64
	Title string
	Uid   int64
	Views int64
}

func init() {

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic("RegisterDriver:" + err.Error())
	}

	// 数据库要自己手动创建。否则报错：[ORM]2019/12/22 10:14:11 register db Ping `default`, Error 1049: Unknown database 'beegoblog'
	//panic: RegisterDataBase:register db Ping `default`, Error 1049: Unknown database 'beegoblog'
	err = orm.RegisterDataBase("default", "mysql", "root:@/beegoblog?charset=utf8")
	if err != nil {
		panic("RegisterDataBase:" + err.Error())
	}

	//// RunSyncdb自动建表。 第一个参数是数据库名字；第二个是是否删除之前的表，重新建表
	orm.RunSyncdb("default", false, true)
	orm.DefaultTimeLoc = time.UTC

	//BlodDb := orm.NewOrm()
	//BlodDb.Using("default")
}

func main() {
	//orm.Debug = true // TODO：放在配置文件中
	//// RunSyncdb自动建表。 第一个参数是数据库名字；第二个是是否删除之前的表，重新建表
	//orm.RunSyncdb("default", false, true)

	beego.SetStaticPath("/down1", "download1")

	beego.Run()
}
