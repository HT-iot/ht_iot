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

/*
{
    "data": [{
        "hospitalname": "ADBD",
        "hospitalzone": "ADBD",
        "hospitalbed": "ADBD",
        "patientname": "ADBD",
        "hospitaldeviceid": "ADBD",
        "channelid": "ADBD",
        "deviceid": "ADBD",
        "puls": 1.0,
        "pressure": 2.0,
        "oxgen": 3.0,
        "position": "ADBD"
    },{
        "hospitalname": "CDE",
        "hospitalzone": "CDE",
        "hospitalbed": "CDE",
        "patientname": "CDE",
        "hospitaldeviceid": "ADBD",
        "channelid": "ADBD",
        "deviceid": "ADBD",
        "puls": 1.0,
        "pressure": 2.0,
        "oxgen": 3.0,
        "position": "ADBD"
    }]
}
*/

type DataTable struct {
/*	Totalpatient	int `json:"totalpatient"`
	Totaluses		int `json:"totaluses"`
	Totalunuses		int `json:"totalunuses"`
	Totalurgent		int `json:"totalurgent"`
	Totalpuls		int `json:"totalpuls"`
	Totaloxgen		int `json:"totaloxgen"`
	Totalpress		int `json:"totalpress"`
*/	Data 			[]PatientInfo `json:"data"`
}

type Warnpara struct{
	Pulsmin             float64 `json:"pulsmin"`
	Pulsmax             float64 `json:"pulsmax"`
	Oxgenmin            float64 `json:"oxgenmin"`
	Oxgenmax            float64 `json:"oxgenmax"`
	Pressurelowmin      float64 `json:"pressurelowmin"`
	Pressurelowmax      float64 `json:"pressurelowmax"`
	Pressurehighmin     float64 `json:"pressurehighmin"`
	Pressurehighmax     float64 `json:"pressurehighmax"`
	Monitoraddress     	string `json:"monitoraddress"`
	Monitorradius     	float64 `json:"monitorradius"`
	MonitorLongitude   	float64 `json:"monitorlongitude"`
	MonitorLatitude     float64  `json:"monitorlatitude"`	
}

type PatientInfo struct {
	Runstatus	     string  `json:"runstatus"`
	Hospitalname     string  `json:"hospitalname"`
	Hospitalzone     string  `json:"hospitalzone"`
	Hospitalbed      string  `json:"hospitalbed"`
	Patientname      string  `json:"patientname"`
	Hospitaldeviceid string  `json:"hospitaldeviceid"`
	Channelid        string  `json:"channelid"`
	Deviceid         string  `json:"deviceid"`
	Puls             float64 `json:"puls"`
	Oxgen            float64 `json:"oxgen"`
	Pressurelow      float64 `json:"pressurelow"`
	Pressurehigh     float64 `json:"pressurehigh"`
	Longitude        float64 `json:"longitude"`
	Latitude         float64 `json:"latitude"`
	Ops 			 string  `json:"ops"`	
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
//need pressure low
		if len(msg) != 0 {
			(*p).Pressurehigh = msg[0].V
			log.Debug(PressureString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Pressurehigh = 0
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
