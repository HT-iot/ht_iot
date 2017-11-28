package models

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/gocql/gocql"
)

/*
var tables []string = []string{
	`CREATE TABLE IF NOT EXISTS messages_by_channel (
		channel timeuuid,
		id timeuuid,
		publisher text,
		protocol text,
		bn text,
		bt double,
		bu text,
		bv double,
		bs double,
		bver int,
		n text,
		u text,
		v double,
		vs text,
		vb boolean,
		vd text,
		s double,
		t double,
		ut double,
		l text,
		PRIMARY KEY ((channel), id)
	) WITH default_time_to_live = 86400`,
}
*/

func GetMsg() {
	log := logs.GetBeeLogger()
	log.Info("Get msg")
	return

}
