package controllers

import (
	"ht_iot/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MgrController struct {
	beego.Controller
}

func (c *MgrController) GetUsers() {
	l := logs.GetBeeLogger()
	l.Debug("get request")

	user, err := models.GetUers()
	if err != nil {
		l.Debug("DB query failure")
	}

	l.Info("user talbe, email is " + user.Email + " password is " + user.Password)
	c.ServeJSON()
}

func (c *MgrController) GetClients() {
	l := logs.GetBeeLogger()
	l.Debug("get request")

	client, err := models.GetClientByUser()
	if err != nil {
		l.Debug("DB query failure")
	}

	for i := 0; i < len(client); i++ {
		l.Info("client " + client[i].Name + client[i].ID.String() + client[i].AccessKey)
	}

}
