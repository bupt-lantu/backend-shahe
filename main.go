package main

import (
	_ "bupt_tour/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}

	//文档链接
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	//log设置
	logs.Reset()
	logs.SetLevel(logs.LevelTrace)
	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/tour.log"}`)

	beego.Run()
}
