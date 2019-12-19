package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"net/url"
)

type LoginController struct {
	beego.Controller
}

const (
	cookieUserName = "uname1"
	cookieUserPwd  = "upwd1"
)

// 进入登陆页面
func (c *LoginController) Get() {
	c.TplName = "login.html"
}

// 发起登陆请求
func (this *LoginController) Post() {
	//str := fmt.Sprint(c.Input())
	//c.Ctx.WriteString(str)

	userName := this.Input().Get("userName")
	userPwd := this.Input().Get("userPwd")
	logs.Debug("enter Post login, %s,%s", userName, userPwd)

	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("username") == userName &&
		beego.AppConfig.String("userpwd") == userPwd {
		maxAge := 50
		if autoLogin {
			maxAge = 86400 * 7
		}
		logs.Info("name and pwd success")

		this.Ctx.SetCookie(cookieUserName, userName, maxAge, "/")
		this.Ctx.SetCookie(cookieUserPwd, userPwd, maxAge, "/")
	}
	this.Redirect("/", 301) // Router.go中定义了 "/"是访问首页(home页)
	return
}

func checkAccount(ctx *context.Context, input url.Values) bool {

	cookie, err := ctx.Request.Cookie(cookieUserName)
	if err != nil {
		fmt.Println("checkAccount 1 error", err)
		return false
	}
	userName := cookie.Value
	cookie, err = ctx.Request.Cookie(cookieUserPwd)
	if err != nil {
		fmt.Println("checkAccount 2 error", err)
		return false
	}
	pwd := cookie.Value
	fmt.Println("userName:", userName, ",pwd:", pwd)

	if beego.AppConfig.String("username") == userName &&
		beego.AppConfig.String("userpwd") == pwd {
		fmt.Println("checkAccount true")
		return true
	}
	fmt.Println("checkAccount false")
	return false
}

type ExitLoginController struct {
	beego.Controller
}

// cookie知识整理:
// 1. cookie如果设置了maxage=0,则页面关闭或者发生跳转等，此cookie会被浏览器删除
// 2。cookie如果设置了maxage=0,页面只要不关闭，发请求的时候，在后端会一直获取到cookie,即浏览器会将cookie携带过去 ==> 错误的
// 3。在后端 设置cookie maxage=-1(只要是负数)，就是让浏览器强制删除cookie
func (this *ExitLoginController) Get() {
	//cookie, err := this.Ctx.Request.Cookie(cookieUserName)
	//if err != nil {
	//	logs.Debug("ExitLoginController err:", err) //http: named cookie not present
	//	this.Ctx.WriteString(err.Error())
	//	return
	//}
	//logs.Debug("ExitLoginController coookie value:", cookie.Value)
	//logs.Debug("ExitLoginController coookie expire:", cookie.Expires)

	this.Ctx.SetCookie(cookieUserName, "", -1, "/")
	this.Ctx.SetCookie(cookieUserPwd, "", -1, "/")
	this.Redirect("/", 301)
	return
}
