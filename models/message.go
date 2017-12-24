package models

import (
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
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

type Message struct {
	Channel   string
	Id        string
	Publisher string
	Protocol  string
	Bn        string `json:"basename,omitempty"`
	Bt        float64
	Bu        string
	Bv        float64
	Bs        float64
	Bver      int
	N         string
	U         string
	V         float64
	Vs        string
	Vb        bool
	Vd        string
	S         float64
	T         float64
	Ut        float64
	L         string
}

type PatientInfo struct {
	Hospitalname     string
	Hospitalzone     string
	Hospitalbed      int
	Patientname      string
	Hospitaldeviceid int
	Channelid        string
	Deviceid         string

	Puls     float64
	Pressure float64
	Oxgen    float64
	Position float64
}

const (
	PulsString     = "puls"
	PressureString = "pressure"
	OxgenString    = "oxgen"
)

/*MAX_MSG define to maxmim number of message for display */
const MaxNum uint = 1000

func GetMsg(p *PatientInfo) error {
	log := logs.GetBeeLogger()
	log.Debug("Get Message for %s", p.Deviceid)
	var msg []Message

	/* only get the last on message */
	sel := qb.Select("messages_by_channel").Where(qb.Eq("n")).Limit(1).AllowFiltering()
	stmt, names := sel.ToCql()

	{ /*get the PulsString*/
		n := p.Deviceid + ":" + PulsString
		q := gocqlx.Query(SessionMsg.Query(stmt), names).BindMap(qb.M{"n": n})
		defer q.Release()

		if err := gocqlx.Select(&msg, q.Query); err != nil {
			log.Error("GetMsg: select Err:" + err.Error())
		}

		if len(msg) != 0 {
			(*p).Puls = msg[0].V
			log.Debug(PulsString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Puls = 0
			log.Error("Can't find in DB, set the value to zero")
		}
	}
	{ /*get the PressureString*/
		n := p.Deviceid + ":" + PressureString
		q := gocqlx.Query(SessionMsg.Query(stmt), names).BindMap(qb.M{"n": n})
		defer q.Release()

		if err := gocqlx.Select(&msg, q.Query); err != nil {
			log.Error("GetMsg: select Err:" + err.Error())
		}

		if len(msg) != 0 {
			(*p).Pressure = msg[0].V
			log.Debug(PressureString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Pressure = 0
			log.Error("Can't find in DB, set the value to zero")
		}
	}
	{ /*get the OxgenString*/
		n := p.Deviceid + ":" + OxgenString
		q := gocqlx.Query(SessionMsg.Query(stmt), names).BindMap(qb.M{"n": n})
		defer q.Release()

		if err := gocqlx.Select(&msg, q.Query); err != nil {
			log.Error("GetMsg: select Err:" + err.Error())
		}

		if len(msg) != 0 {
			(*p).Oxgen = msg[0].V
			log.Debug(OxgenString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Oxgen = 0
			log.Error("Can't find in DB, set the value to zero")
		}
	}
	return nil
}
