package main

import (
  "fmt"
  _  "github.com/bupt-lantu/backend-shahe/routers"

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
  orm.Debug = true
  orm.SetMaxOpenConns("default", 200)
  orm.SetMaxIdleConns("default", 300)

  fmt.Println(beego.BConfig.WebConfig.Session.SessionOn)

  //文档链接
  beego.BConfig.WebConfig.DirectoryIndex = true
  beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

  //log设置
  logs.Reset()
  logs.SetLogger(logs.AdapterFile, `{"filename":"logs/tour.log"}`)

  beego.Run()
}
