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
	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/users", &controllers.MgrController{}, "get:GetUsers")
	beego.Router("/clients", &controllers.MgrController{}, "get:GetClients")

	beego.Router("/status", &controllers.StatusController{})
	beego.Router("/api/status", &controllers.StatusController{}, "get:GetStatus;post:PostStatus")
	beego.Router("/warn/status", &controllers.StatusController{}, "get:GetWarnStatus;post:PostWarnStatus")
	beego.Router("/info/status", &controllers.StatusController{}, "get:GetInfo;post:PostInfo")
	beego.Router("/day/status", &controllers.StatusController{}, "get:GetOneData")
//	beego.Router("/day/status", &controllers.StatusController{}, "get:GetOne;post:GetFive")
	
	beego.Router("/pconfig", &controllers.PconfigController{})
	beego.Router("/api/pconfig", &controllers.PconfigController{}, "get:GetPat;post:PostPat")
	beego.Router("/line/pconfig", &controllers.PconfigController{}, "get:GetLine;post:PostLine")
	beego.Router("/device/pconfig", &controllers.PconfigController{}, "post:GetDevInfo")

	beego.Router("/deviceinfo", &controllers.DeviceinfoController{})
	beego.Router("/api/deviceinfo", &controllers.DeviceinfoController{}, "get:GetModal;post:PostModal")
	beego.Router("/merge/deviceinfo", &controllers.DeviceinfoController{}, "post:PostMerge")

}
