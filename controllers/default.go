package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["TrueCondition"] = true
	c.Data["falseCondition"] = false
	c.Data["array"] = []int{1, 2, 3, 4, 5}
	c.Data["myhtml"] = "<div>hahah</div>"
	c.Data["testVar"] = "test tpl var"

	type User struct {
		Name string
		Id   int
	}
	c.Data["user"] = User{Name: "john", Id: 100}

	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.TplName = "TestPost.tpl"
}

func (c *UserController) Get() {
	c.Ctx.WriteString("this is from userController by get")
}

func (c *UserController) GetUser() {
	c.Ctx.WriteString("this is mapping controller ")
}

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Ctx.WriteString("hello world")
}
