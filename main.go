package main

import (
	"ht_iot/models"
	_ "ht_iot/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {

	log := logs.GetBeeLogger()

	err := models.InitDB()
	if err != nil {
		log.Critical(err.Error())
	}

	log.Info("connected to cassandra ...")

	beego.Run()
}
