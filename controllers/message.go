package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MsgController struct {
	beego.Controller
}

func (c *MsgController) GetMsg() {
	l := logs.GetBeeLogger()
	l.Debug("get request")
}
