package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "myBeego/routers"
)

//func init() {
//	models.RegisterDB()
//}

var BlodDb orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@/beegoBlog?charset=utf8")

	BlodDb := orm.NewOrm()
	BlodDb.Using("default")
}

func main() {
	//orm.Debug = true // TODO：放在配置文件中
	//// RunSyncdb自动建表。 第一个参数是数据库名字；第二个是是否删除之前的表，重新建表
	//orm.RunSyncdb("default", false, true)

	beego.SetStaticPath("/down1", "download1")

	beego.Run()
}
