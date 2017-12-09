package controllers

import (
	"github.com/astaxie/beego"
)

//MainController ....
type MainController struct {
	beego.Controller
}

var IsLogin, IsStatus, IsPconfig, IsOconfig bool

//Get .....
func (c *MainController) Get() {

	c.Data["Website"] = "www.hitech_iot.com"
	c.Data["Email"] = "jason.zhang@hitech_iot.com"
	c.TplName = "home.html"
	//	c.Data["IsLogin"] = checkAccount(c.Ctx)
	IsLogin = checkAccount(c.Ctx)
	c.Data["ISLogin"] = IsLogin

	if !IsLogin {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Data["IsLogin"] = false
		c.Redirect("/login", 302)
		return
	}

	/*
		isExit, _ := c.GetBool("IsLogin")
		isExit2, _ := c.GetInt("exit")

		if isExit && (isExit2 == 9) {
			c.Ctx.SetCookie("uname", "", -1, "/")
			c.Ctx.SetCookie("pwd", "", -1, "/")
			c.Data["IsLogin"] = false
			//		c.Redirect("/", 301)
			return
		}
	*/
}
