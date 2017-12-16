package routers

import (
	"ht_iot/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	l := logs.GetBeeLogger()
	l.Info("init router ...")

	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.MgrController{}, "get:GetUsers")
	beego.Router("/clients", &controllers.MgrController{}, "get:GetClients")
	beego.Router("/msg", &controllers.MsgController{}, "get:GetMsg")
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/pconfig", &controllers.PconfigController{})
	beego.Router("/outhospital", &controllers.OuthospitalController{})
	beego.Router("/deviceinfo", &controllers.DeviceinfoController{})
}
