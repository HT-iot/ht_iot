package models

import (
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"time"
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
	Reporttime		 string  `json:"reporttime"`		
	Bodycapacitance	 float64 `json:"bodycapacitance"`
	Puls             float64 `json:"puls"`
	Oxgen            float64 `json:"oxgen"`
	Pressurelow      float64 `json:"pressurelow"`
	Pressurehigh     float64 `json:"pressurehigh"`
	Longitude        float64 `json:"longitude"`
	Latitude         float64 `json:"latitude"`
	Ops 			 string  `json:"ops"`	
}

//parameter of pulse data
const (
	BodyString = "Body"
	PulseString     = "Pulse"   
	PressureHString = "Presshigh"
	PressureLString = "Presslow"
	OxgenString    = "Oxgen"
	LongitudeString    = "Longitude"
	LatitudeString = "Latitude"
)

/*MAX_MSG define to maxmim number of message for display */
const MaxNum uint = 1000

func GetMsg(p *PatientInfo) error {
	log := logs.GetBeeLogger()
	log.Debug("Get Message for %s", p.Deviceid)
	var msg []Message
//	var pulse_time, oxgen_time, presshigh_time, presslow_time
	var bodycapacitancetime, pulsetime, oxgentime, presshightime,presslowtime,latitudetime,longitudetime float64

	/* only get the last on message */
	sel := qb.Select("messages_by_channel").Where(qb.Eq("n")).Limit(1).AllowFiltering()
	stmt, names := sel.ToCql()

	
	{ /*get the Body Capacitor*/
		n:=p.Deviceid+":"+ BodyString
		q := gocqlx.Query(SessionMsg.Query(stmt), names).BindMap(qb.M{"n": n})
		defer q.Release()

		if err := gocqlx.Select(&msg, q.Query); err != nil {
			log.Error("GetMsg: select Err:" + err.Error())
		}

		if len(msg) != 0 {
			(*p).Bodycapacitance = msg[0].V
			bodycapacitancetime = msg[0].T
			log.Debug(PulseString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Bodycapacitance = 0
			log.Error("Can't find in DB, set the value to zero")
		}
	}
	
	{ /*get the PulsString*/
		n:=p.Deviceid+":"+ PulseString
		q := gocqlx.Query(SessionMsg.Query(stmt), names).BindMap(qb.M{"n": n})
		defer q.Release()

		if err := gocqlx.Select(&msg, q.Query); err != nil {
			log.Error("GetMsg: select Err:" + err.Error())
		}

		if len(msg) != 0 {
			(*p).Puls = msg[0].V
			pulsetime = msg[0].T
			log.Debug(PulseString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Puls = 0
			log.Error("Can't find in DB, set the value to zero")
		}
	}
	{ /*get the Pressure High*/
		n := p.Deviceid + ":" + PressureHString

		q := gocqlx.Query(SessionMsg.Query(stmt), names).BindMap(qb.M{"n": n})
		defer q.Release()

		if err := gocqlx.Select(&msg, q.Query); err != nil {
			log.Error("GetMsg: select Err:" + err.Error())
		}
//need pressure low
		if len(msg) != 0 {
			(*p).Pressurehigh = msg[0].V
			presshightime = msg[0].T
			log.Debug(PressureHString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Pressurehigh = 0
			log.Error("Can't find in DB, set the value to zero")
		}
	}
	{ /*get the Pressure Low*/
		n := p.Deviceid + ":" + PressureLString
		q := gocqlx.Query(SessionMsg.Query(stmt), names).BindMap(qb.M{"n": n})
		defer q.Release()

		if err := gocqlx.Select(&msg, q.Query); err != nil {
			log.Error("GetMsg: select Err:" + err.Error())
		}
//need pressure low
		if len(msg) != 0 {
			(*p).Pressurelow = msg[0].V
			presslowtime = msg[0].T
			log.Debug(PressureLString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Pressurelow = 0
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
			oxgentime = msg[0].T
			log.Debug(OxgenString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Oxgen = 0
			log.Error("Can't find in DB, set the value to zero")
		}
	}
	{ /*get the LatitudeString*/
		n := p.Deviceid + ":" + LatitudeString
		q := gocqlx.Query(SessionMsg.Query(stmt), names).BindMap(qb.M{"n": n})
		defer q.Release()

		if err := gocqlx.Select(&msg, q.Query); err != nil {
			log.Error("GetMsg: select Err:" + err.Error())
		}

		if len(msg) != 0 {
			(*p).Latitude = msg[0].V
			latitudetime = msg[0].T
			log.Debug(LatitudeString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Latitude = 0
			log.Error("Can't find in DB, set the value to zero")
		}
	}
	{ /*get the LongitudeString*/
		n := p.Deviceid + ":" + LongitudeString
		q := gocqlx.Query(SessionMsg.Query(stmt), names).BindMap(qb.M{"n": n})
		defer q.Release()

		if err := gocqlx.Select(&msg, q.Query); err != nil {
			log.Error("GetMsg: select Err:" + err.Error())
		}

		if len(msg) != 0 {
			(*p).Longitude = msg[0].V
			longitudetime = msg[0].T
			log.Debug(LongitudeString + "value is " + strconv.FormatFloat(msg[0].V, 'f', 6, 64))
		} else {
			(*p).Longitude = 0
			log.Error("Can't find in DB, set the value to zero")
		}
	}
	//需要判断最近时间
	t1:=time.Now()
	t2:= time.Unix(int64(pulsetime),0);
	d := t1.Sub(t2).Seconds()
	if ((pulsetime == oxgentime)&&(pulsetime == presshightime)&&(pulsetime==presslowtime)&&(pulsetime == longitudetime)&& (pulsetime == latitudetime)&&(pulsetime == bodycapacitancetime)){
		if (d<=300){
			if((*p).Bodycapacitance>=5){
				(*p).Runstatus = "在线使用"  
				(*p).Reporttime = (*p).Runstatus +"："+ t2.Format("2006-01-02 03:04:05 PM")
				if((*p).Puls<30){
					(*p).Runstatus = "在线危急"  
					(*p).Reporttime = (*p).Runstatus +"："+ t2.Format("2006-01-02 03:04:05 PM")
				}
			}else{
				(*p).Runstatus = "在线未使用"  
				(*p).Reporttime = (*p).Runstatus +"："+ t2.Format("2006-01-02 03:04:05 PM")
			}

		} else{
			(*p).Runstatus = "离线"
			(*p).Reporttime = (*p).Runstatus +"："+ t2.Format("2006-01-02 03:04:05 PM")
//			(*p).Reporttime = (*p).Runstatus +"："+ strconv.FormatInt(int64(d),10) +"秒"
		}
	}else{
		(*p).Runstatus = "数据不完整"  
		(*p).Reporttime = (*p).Runstatus +"："+ t2.Format("2006-01-02 03:04:05 PM")
	}
	return nil
}
