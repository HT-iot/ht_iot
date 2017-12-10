package controllers

import (
	"fmt"
	"ht_iot/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
	isExit := c.GetString("exit")
	fmt.Println("Login isExit =", isExit)
	if isExit == "true" {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
}

func (c *LoginController) Post() {
	c.TplName = "login.html"
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")

	autoLogin := c.Input().Get("autoLogin") == "on"

	p := models.User{Email: uname, Password: pwd}
	q, err := models.GetAllUers(p)

	fmt.Println("q, err =", q, err)

	if len(q) > 0 {
		fmt.Println("register pwd")
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		IsLogin = true
		c.Data["ISLogin"] = IsLogin
		fmt.Println("IsLogin =", IsLogin)

		//, IsStatus, IsPconfig, IsOconfig
		c.TplName = "home.html"
	} else {
		IsLogin = false
		c.Data["ISLogin"] = IsLogin
		c.Redirect("/login", 301)
	}
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	p := models.User{Email: uname, Password: pwd}
	q, err := models.GetAllUers(p)

	//	fmt.Println("q, err =", q, err)

	if len(q) > 0 {
		IsLogin = true
		return true
	} else {
		IsLogin = false
		return false

	}
}
