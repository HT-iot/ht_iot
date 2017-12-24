package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"ht_iot/models"
	"io"

	"github.com/astaxie/beego/logs"

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
		c.Ctx.SetCookie("name", "", -1, "/")
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
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("name", uname, maxAge, "/")
		h := sha1.New()
		io.WriteString(h, uname+pwd)
		pwd = base64.StdEncoding.EncodeToString(h.Sum(nil))
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		/* TODO add session support
		c.SetSession(uname, pwd)
		*/
		IsLogin = true
		c.Data["ISLogin"] = IsLogin

		c.TplName = "home.html"
	} else {
		IsLogin = false
		c.Data["ISLogin"] = IsLogin
		c.Redirect("/login", 301)
	}
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("name")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	/*
		TODO add the session support in the furture to reduce query user table

		sid := ctx.Input.CruSession.Get(uname)
		logs.Debug("get session", sid)
	*/

	p := models.User{Email: uname, Password: ""}
	q, err := models.GetAllUers(p)
	//fmt.Println("q, err =", q, err)

	if len(q) < 1 {
		logs.Debug("Query user table failure")
		IsLogin = false
		return false
	}

	h := sha1.New()
	io.WriteString(h, q[0].Email+q[0].Password)
	if pwd == base64.StdEncoding.EncodeToString(h.Sum(nil)) {
		logs.Debug("compare pwd success")
		IsLogin = true
		return true
	}

	logs.Debug("Cookie password check failure")
	IsLogin = false
	return false

}
