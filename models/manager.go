package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
)

var tables = []string{
	`CREATE TABLE IF NOT EXISTS manager.hospital_by_patient (
		hospitalname text,
		id timeuuid,
		hospitalzone text,
		patientname text,
		hospitaldeviceid text,
		inhospital boolean,
		channelid text,
		deviceid text,
		hospitalbed text,
		hospitalid text,
		patiententrtime timestamp,
		patientexittime timestamp,
		patientid text,
		patientsex text,
		meta map<text, text>,
		PRIMARY KEY ((hospitalname),deviceid, channelid)
		)`,

	`CREATE TABLE IF NOT EXISTS manager.device_info (
		hospitalname text,
		manufacturer  text,
		modelnumber text,
		deviceid text,
		firmwareVersion text,
		reboot boolean,
		factoryreset boolean,
		availablepowersource int,
		powersourcevoltage int,
		powersourcesurrent int,
		batterylevel int,
		memoryfree int,
		errorcode int,
		reseterrorcode text,
		currenttime timestamp,
		utcoffset text,
		timezone text,
		supportedbindingmodes text,
		devicetype text,
		hardwareversion text,
		softwareversion text,
		hospitaldeviceid text,
		channelid text,
		used boolean,
		PRIMARY KEY ((hospitalname),deviceid,channelid)
		)`,
	 
	`CREATE TABLE IF NOT EXISTS manager.hospital_config (
		hospitalname text,
		hospitalzone text,
		pulsmin double,   
		pulsmax double,   
		oxgenmin double,    
		oxgenmax double,     
		pressurelowmin double,
		pressurelowmax double,
		pressurehighmin double,
		pressurehighmax double,
		monitoraddress text, 
		monitorradius double,
		latitude double,
		longitude double,
		ops text,
		PRIMARY KEY (hospitalname,hospitalzone)
	)`,
}

/*User : The Admin user information */
type User struct {
	Email    string
	Password string
}

/*ClientsByUser : The Device/App information */
type ClientsByUser struct {
	User      string
	ID        gocql.UUID
	Type      string
	Name      string
	AccessKey string
	Meta      map[string]string
}

/*ChannelsByUser : The Device/Channel information */
type ChannelsByUser struct {
	User      string
	ID        gocql.UUID
	Name      string
	Connected []string
}

type HospitalInfoTable struct {
	Data HospitalInfoConfig `json:"data"`
}

/*医院病区监控的阈值信息*/
type HospitalInfoConfig struct {
	Hospitalname  	string 		`json:"hospitalname"`
	Hospitalzone 	string 		`json:"hospitalzone"`
	Pulsmin 		float64   	`json:"pulsmin"`
	Pulsmax 		float64   	`json:"pulsmax"`
	Oxgenmin 		float64    	`json:"oxgenmin"`
	Oxgenmax 		float64     `json:"oxgenmax"`
	Pressurelowmin 	float64 	`json:"pressurelowmin"`
	Pressurelowmax 	float64 	`json:"pressurelowmax"`
	Pressurehighmin float64		`json:"pressurehighmin"`
	Pressurehighmax float64 	`json:"pressurehighmax"`
	Monitoraddress 	string  	`json:"monitoraddress"`
	Monitorradius 	float64 	`json:"monitorradius"`
	Latitude 		float64 	`json:"latitude"`
	Longitude 		float64 	`json:"longitude"`
	Ops 			string 		`json:"ops"`
}


type HospitalTable struct {
	Data []HospitalPatientInfo `json:"data"`
}

/*HospitalPatientInfo : The Hospital/Patient information */
type HospitalPatientInfo struct {
	Id               gocql.UUID
	Hospitalname     string
	Hospitalzone     string
	Patientname      string
	Hospitaldeviceid string
	Inhospital       bool
	Channelid        string
	Deviceid         string
	Hospitalbed      string
	Hospitalid       string
	Patiententrtime  time.Time
	Patientexittime  time.Time
	Patientid        string
	Patientsex       string
	Meta             map[string]string
}

/*DeviceInfo : The Device information */
type DeviceTable struct {
	Data []DeviceInfo `json:"data"`
}

type DeviceInfo struct {
	Hospitalname          string
	Manufacturer          string
	Modelnumber           string
	Deviceid              string
	Firmwareversion       string
	Reboot                bool
	Factoryreset          bool
	Availablepowersource  int
	Powersourcevoltage    int
	Powersourcesurrent    int
	Batterylevel          int
	Memoryfree            int
	Errorcode             int
	Reseterrorcode        string
	Currenttime           time.Time
	Utcoffset             string
	Timezone              string
	Supportedbindingmodes string
	Devicetype            string
	Hardwareversion       string
	Softwareversion       string
	Hospitaldeviceid      string
	Channelid             string
	Used                  bool
}

/*GetUsers  get data... */
func GetUers() (User, error) {
	log := logs.GetBeeLogger()

	stmt, names := qb.Select("users").ToCql()
	log.Info("stmt is " + stmt)
	q := gocqlx.Query(SessionMgr.Query(stmt), names)
	defer q.Release()

	var user User
	var err error
	if err = gocqlx.Get(&user, q.Query); err != nil {
		log.Critical("GetUers:" + err.Error())
	}

	return user, err
}

/*GetUsers  get all data...*/
func GetAllUers(p User) ([]User, error) {
	log := logs.GetBeeLogger()
	sel := qb.Select("users").Where(qb.Eq("email")).Limit(100).AllowFiltering()

	/*
	if p.Password != "" {
		sel.Where(qb.Eq("password"))
	}
	*/
	stmt, names := sel.ToCql()

	//fmt.Println("stmt =", stmt)
	//fmt.Println("names =", names)

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&p)
	//fmt.Println("q.Query=  ", q.Query)
	defer q.Release()

	var people []User
	var err error
	if err = gocqlx.Select(&people, q.Query); err != nil {
		log.Critical("GetAllUers:" + err.Error())
	}

	return people, err
}

func SaveUers(u *User) error {
	log := logs.GetBeeLogger()
	log.Info("start SaveUser")

	log.Info("start Insert admin information")
	stmt, names := qb.Insert("users").Columns("email", "password").ToCql()

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(u)

	if err := q.ExecRelease(); err != nil {
		log.Critical("select:" + err.Error())
		return err
	}
	return nil
}

/*GetClientByUser  get data... */
func GetClientByUser() ([]ClientsByUser, error) {
	log := logs.GetBeeLogger()
	log.Info("start the GetClientByUser")

	stmt, names := qb.Select("clients_by_user").Limit(200).ToCql()

	log.Info("stmt is " + stmt)
	//	fmt.Println(names)
	q := gocqlx.Query(SessionMgr.Query(stmt), names)
	defer q.Release()
	var clients []ClientsByUser
	var err error
	if err = gocqlx.Select(&clients, q.Query); err != nil {
		log.Critical("select:" + err.Error())
	}

	return clients, err
}

/*InsertPatient Insert Patient to Cassandra
func GetPatient(h HospitalPatientInfo) (*[]HospitalPatientInfo, error) {
	// Insert with query parameters bound from struct.

	log := logs.GetBeeLogger()
	log.Info("Get Patient information")

	sel := qb.Select("hospital_by_patient").Where(qb.Eq("hospitalname")).AllowFiltering()

		if len(h.Hospitalzone) > 0 {
			sel.Where(qb.Eq("hospitalzone"))
		}

		if len(h.Patientname) > 0 {
			sel.Where(qb.Eq("patientname"))
		}

	stmt, names := sel.ToCql()

	fmt.Println("stmt =", stmt)
	fmt.Println("names =", names)

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)
	fmt.Println("q.Query=  ", q.Query)
	defer q.Release()

	var Patient []HospitalPatientInfo
	if err := gocqlx.Select(Patient, q.Query); err != nil {
		fmt.Println("select Err:", err)
	}
	return &Patient, nil
}
*/
/*InsertPatient Insert Patient to Cassandra */
func InsertPatient(h HospitalPatientInfo) error {
	// Insert with query parameters bound from struct.
	log := logs.GetBeeLogger()
	log.Info("Insert Patient information")

	stmt, names := qb.Insert("hospital_by_patient").Columns("id", "hospitalname", "hospitalzone",
		"patientname", "hospitaldeviceid", "inhospital", "channelid",
		"deviceid", "hospitalbed", "hospitalid", "patiententrtime",
		"patientid", "patientsex", "meta").
		ToCql()

		//	fmt.Println("stmt =", stmt)
		//	fmt.Println("h =", h)
	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)
	//	fmt.Println("q =", q)

	if err := q.ExecRelease(); err != nil {
		log.Critical("InsertPatient: select:" + err.Error())
		return err
	}
	return nil
}

func GetAllPatient() ([]HospitalPatientInfo, error) {
	log := logs.GetBeeLogger()
	log.Debug("Get ALL Patient")

	sel := qb.Select("hospital_by_patient").Limit(MaxNum).AllowFiltering()
	stmt, names := sel.ToCql()

	q := gocqlx.Query(SessionMgr.Query(stmt), names)
	defer q.Release()

	var patient []HospitalPatientInfo
	if err := gocqlx.Select(&patient, q.Query); err != nil {
		log.Error("GetAllPatient: select Err:" + err.Error())
	}
	return patient, nil
}

func GetPatient(h HospitalPatientInfo) ([]HospitalPatientInfo, error) {
	// Insert with query parameters bound from struct.

	log := logs.GetBeeLogger()
	log.Debug("Get Patient information")

	sel := qb.Select("hospital_by_patient").Limit(MaxNum).AllowFiltering()
	//sel := qb.Select("hospital_by_patient").Where(qb.Eq("hospitalname")).Limit(MaxNum).AllowFiltering()

	if len(h.Hospitalname) != 0 {
		sel.Where(qb.Eq("hospitalname"))
	}
	if len(h.Hospitalzone) != 0 {
		sel.Where(qb.Eq("hospitalzone"))
	}
	if len(h.Hospitalbed) != 0 {
		sel.Where(qb.Eq("hospitalbed"))
	}
	if len(h.Patientname) != 0 {
		sel.Where(qb.Eq("patientname"))
	}

	if len(h.Deviceid) != 0 {
		sel.Where(qb.Eq("deviceid"))
	}

	stmt, names := sel.ToCql()
	logs.Debug(stmt, names)
	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)
	defer q.Release()

	var patient []HospitalPatientInfo
	if err := gocqlx.Select(&patient, q.Query); err != nil {
		log.Debug("select Err:", err.Error())
	}
	return patient, nil
}

func GetPaientIDs(name, hid string) (ch, de string, hv bool) {

	log := logs.GetBeeLogger()
	log.Info("Patient to Device mapping")

	err := SessionMgr.Query(`SELECT Channelid, Deviceid from device_info WHERE Hospitalname = ? AND hospitaldeviceid=? AND Used !=? LIMIT 1 ALLOW FILTERING`, name,hid,true).Scan(&ch,&de)

	if(err!=nil){
		log.Debug("select Err:", err.Error())
		return "","", false
	}
	fmt.Println(ch,de)
	return ch,de,true


/*
	d := DeviceInfo{Hospitalname: name, Hospitaldeviceid: hid}
	sel := qb.Select("device_info").Where(qb.Eq("hospitalname")).AllowFiltering()
	//	fmt.Println("d:", d)

	if len(hid) != 0 {
		sel.Where(qb.Eq("hospitaldeviceid"))
	}

	stmt, names := sel.ToCql()
	//	fmt.Println("stmt,names:", stmt, names)
	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&d)
	defer q.Release()

	var device []DeviceInfo
	if err := gocqlx.Select(&device, q.Query); err != nil {
		fmt.Println("select Err:", err)
		return "", "", false
	}
	if len(device) > 0 && (!device[0].Used) {
		ch = device[0].Channelid
		de = device[0].Deviceid
		device[0].Used = true
		UpdateDeviceItem(device[0])
		return ch, de, true
	}
	return "", "", false
*/	
}

func UpdatePatient(h HospitalPatientInfo) bool {
	// // Easy update with all parameters bound from struct.

	//		p.Email = append(p.Email, "patricia1.citzen@gocqlx_test.com")
	stmt, names := qb.Update("hospital_by_patient").
		Set("hospitalzone", "hospitalbed", "patientname", "patientsex", "patientid", "hospitaldeviceid").
		Where(qb.Eq("hospitalname"), qb.Eq("deviceid"), qb.Eq("channelid")).
		ToCql()
	fmt.Println("Update h=", h)
	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)
	fmt.Println("q=", q)
	if err := q.ExecRelease(); err != nil {
		fmt.Println("err=", err)
		return false
	}
	return true
}


//Update hospital db data  	err := models.UpdateWarnInfo(h)
func UpdateWarnInfo(h HospitalInfoConfig) bool {
	// // Easy update with all parameters bound from struct.
	log := logs.GetBeeLogger()
	//		p.Email = append(p.Email, "patricia1.citzen@gocqlx_test.com")
	sel:= qb.Update("hospital_config").
		Set("pulsmin", "pulsmax", "oxgenmin", "oxgenmax", "pressurelowmin", "pressurelowmax",
			"pressurehighmin", "pressurehighmax","monitoraddress", "monitorradius").
		Where(qb.Eq("hospitalname"))
	
		if len("hospitalzone") != 0 {
			sel.Where(qb.Eq("hospitalzone"))
		}

		stmt, names := sel.ToCql()

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)

	if err := q.ExecRelease(); err != nil {
		log.Error("GetHospital config: select Err:")
		return false
	}
	return true
}


func InsertWarnInfo(h HospitalInfoConfig) error {
	// Insert with query parameters bound from struct.
	log := logs.GetBeeLogger()
	log.Info("Insert hospital config information")

	stmt, names := qb.Insert("hospital_config").Columns("hospitalname","hospitalzone",
		"pulsmin", "pulsmax", "oxgenmin", "oxgenmax", "pressurelowmin",
		"pressurelowmax","pressurehighmin", "pressurehighmax","monitoraddress", "monitorradius").
		ToCql()

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)
	
	if err := q.ExecRelease(); err != nil {
		log.Critical("Insert hospital config information: select:" + err.Error())
		return err
	}
	return nil
}

func GetWarnInfo(hospitalname,hospitalzone string ) (HospitalInfoConfig, error) {
	log := logs.GetBeeLogger()
	log.Debug("Get Hospital defined warning threshold")

	var hinfo []HospitalInfoConfig
	d := HospitalInfoConfig{Hospitalname:hospitalname, Hospitalzone: hospitalzone}
	
	if (len(hospitalname)>0){
		
		sel := qb.Select("hospital_config").Where(qb.Eq("hospitalname")).Limit(1).AllowFiltering()
		if len("hospitalzone") != 0 {
			sel.Where(qb.Eq("hospitalzone"))
		}
	
		stmt, names := sel.ToCql()
		q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&d)
		defer q.Release()

		err := gocqlx.Select(&hinfo, q.Query)
		
		if ((err != nil)||(len(hinfo)==0)) {
//			log.Error("GetHospital config: select Err:")
			//use default value
			d.Hospitalname = hospitalname
			d.Hospitalzone = hospitalzone
			d.Pulsmin = 40
			d.Pulsmax = 120
			d.Oxgenmin = 90
			d.Oxgenmax = 110
			d.Pressurelowmin = 50
			d.Pressurelowmax = 120
			d.Pressurehighmin = 70
			d.Pressurehighmax = 180
			d.Monitoraddress = "深圳市眼科医院"
			d.Monitorradius = 999999
			return d, nil
		} else
		{
			return hinfo[0],nil
		} 
	}else
	{
		d.Pulsmin = 40
		d.Pulsmax = 120
		d.Oxgenmin = 90
		d.Oxgenmax = 110
		d.Pressurelowmin = 50
		d.Pressurelowmax = 120
		d.Pressurehighmin = 70
		d.Pressurehighmax = 180
		d.Monitoraddress = "深圳市眼科医院"
		d.Monitorradius = 999999
		return d, nil
	}
}


func GetDevfromPatient(h HospitalPatientInfo) (HospitalPatientInfo, error) {
	// Check the device & channel ID from Patient information.

	log := logs.GetBeeLogger()
	log.Debug("Get Patient information")

	sel := qb.Select("hospital_by_patient").Where(qb.Eq("hospitalname")).AllowFiltering()
	if len(h.Hospitalzone) != 0 {
		sel.Where(qb.Eq("hospitalzone"))
	}
	if len(h.Hospitalbed) != 0 {
		sel.Where(qb.Eq("hospitalbed"))
	}
	if len(h.Hospitaldeviceid) != 0 {
		sel.Where(qb.Eq("hospitaldeviceid"))
	}
	if len(h.Patientname) != 0 {
		sel.Where(qb.Eq("patientname"))
	}
//	fmt.Println("sel=",sel);
	stmt, names := sel.ToCql()
//	logs.Debug(stmt, names)
	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)
	defer q.Release()
//	fmt.Println("q=",q);
	var patient []HospitalPatientInfo
	if err := gocqlx.Select(&patient, q.Query); err != nil {
		log.Debug("select Err:", err.Error())
		return patient[0], err
	} 
//	fmt.Println("h=",patient);
	return patient[0], nil
}



func GetDevKey(d string) (string) {
	// Check the device & channel ID from Patient information.

	log := logs.GetBeeLogger()
	log.Debug("Get Device access key from Deviceid")	

	var Key string;

	err := SessionMgr.Query(`SELECT access_key from clients_by_user WHERE id = ? LIMIT 1 ALLOW FILTERING`, d).Scan(&Key)

	if(err!=nil){
		log.Debug("select Err:", err.Error())
	}
	fmt.Println(Key)
	return Key
}

