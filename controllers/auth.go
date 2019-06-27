package controllers

import (
	"github.com/astaxie/beego"
)

// AuthController operations for Auth
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @Title login
// @Description 登录
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 string ""
// @Failure 403 body is empty
// @router / [post]
func (c *AuthController) Post() {
	if c.GetString("username") == "super" && c.GetString("password") == "super" {
		c.SetSession("id", int(12));
		c.Data["json"] = "登录成功"
	} else {
		c.Data["json"] = "用户名或密码错误！"
		c.Ctx.ResponseWriter.WriteHeader(400)
	}
	c.ServeJSON()
}

// Get ...
// @Title 检测登录状态
// @Description
// @Success 200 string ""
// @router / [get]
func (c *AuthController) Get() {
	_, ok := c.GetSession("id").(int)
	if ok {
		c.Ctx.WriteString("{\"data\":\"ok\"}")
	} else {
		c.Ctx.WriteString("{\"data\":\"未登录\"}")
	}
}
