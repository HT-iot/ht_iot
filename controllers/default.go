package controllers

import (
	"github.com/astaxie/beego"
)

//MainController ....
type MainController struct {
	beego.Controller
}

//Get .....
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "home.html"

	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Data["IsLogin"] = false
		c.Redirect("/", 301)
		return
	} else {
		c.Data["IsLogin"] = checkAccount(c.Ctx)
	}

}
