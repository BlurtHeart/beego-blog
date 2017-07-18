package main

import (
	"beego-blog/models"
	_ "beego-blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	maxIdle := 5
	maxConn := 5
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./foo.db", maxIdle, maxConn)
	orm.RunSyncdb("default", false, true)
	models.InsertRoles()
}

func main() {
	beego.Run()
}
