package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "404.tpl"
}

func (c *ErrorController) Error400() {
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	fmt.Println(c.Data["json"])
	c.ServeJSON()
}
