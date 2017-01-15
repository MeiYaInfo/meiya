package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
	_ "meiya/server/models"
	_ "meiya/server/routers"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
	var dbUser = beego.AppConfig.String("dbuser")
	var dbPasswd = beego.AppConfig.String("dbpasswd")
	var dbName = beego.AppConfig.String("dbname")
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	conn := dbUser + ":" + dbPasswd + "@/" + dbName + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn /* "root:root@/company?charset=utf8"*/)
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.SetStaticPath("static", "static")
	beego.Run()
}
