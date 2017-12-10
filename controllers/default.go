package controllers

import (
	"fmt"

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
	IsLogin = checkAccount(c.Ctx)
	c.Data["ISLogin"] = IsLogin

	isExit := c.GetString("exit")
	fmt.Println("Main isExit =", isExit)

	if isExit == "true" {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		IsLogin = false
		c.Data["ISLogin"] = IsLogin
		c.Redirect("/", 302)
		return
	}

}
