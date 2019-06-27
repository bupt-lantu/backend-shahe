package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Ctx.WriteString("page not found")
	c.Ctx.ResponseWriter.Status = 404
}

func (c *ErrorController) Error400() {
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.ServeJSON()
}

func (c *ErrorController) Error403() {
	c.ServeJSON()
}
