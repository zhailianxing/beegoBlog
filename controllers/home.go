package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {
	c.Data["isLogined"] = checkAccount(c.Ctx, c.Input())
	logs.Info("islogined:", c.Data["isLogined"])
	c.TplName = "home.html"
}
