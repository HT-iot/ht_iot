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
		hospitaldeviceid int,
		inhospital boolean,
		channelid text,
		deviceid text,
		hospitalbed int,
		hospitalid text,
		patiententrtime timestamp,
		patientexittime timestamp,
		patientid text,
		patientsex text,
		meta map<text, text>,
		PRIMARY KEY (hospitalname,id, hospitalzone, patientname, hospitaldeviceid)
		)`,

	`CREATE TABLE IF NOT EXISTS manager.device_info (
		hospitalname text,
		id timeuuid,
		manufacturer  text,
		modelnumber text,
		Deviceid text,
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
		hospitaldeviceid int,
		channelid text,
		PRIMARY KEY (hospitalname,hospitaldeviceid,deviceid,channelid)
		)`,

	/* These Table will be created by manager application
	`CREATE TABLE IF NOT EXISTS users (
		email text,
		password text,
		PRIMARY KEY (email)
	)`,
	`CREATE TABLE IF NOT EXISTS clients_by_user (
		user text,
		id timeuuid,
		type text,
		name text,
		access_key text,
		meta map<text, text>,
		PRIMARY KEY ((user), id)
	)`,
	`CREATE TABLE IF NOT EXISTS channels_by_user (
		user text,
		id timeuuid,
		name text,
		connected set<text>,
		PRIMARY KEY ((user), id)
	)`,
	`CREATE MATERIALIZED VIEW IF NOT EXISTS clients_by_channel
		AS SELECT user, id, connected FROM channels_by_user
		WHERE id IS NOT NULL
		PRIMARY KEY (id, user)
	`,
	*/
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

/*HospitalPatientInfo : The Hospital/Patient information */
type HospitalPatientInfo struct {
	Id               gocql.UUID
	Hospitalname     string
	Hospitalzone     string
	Patientname      string
	Hospitaldeviceid int
	Inhospital       bool
	Channelid        string
	Deviceid         string
	Hospitalbed      int
	Hospitalid       string
	Patiententrtime  time.Time
	Patientexittime  time.Time
	Patientid        string
	Patientsex       string
	Meta             map[string]string
}

/*DeviceInfo : The Device information */
type DeviceInfo struct {
	Hospitalname          string
	Id                    gocql.UUID
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
	Hospitaldeviceid      int
	Channelid             string
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
		log.Critical("select:" + err.Error())
	}

	return user, err
}

/*GetUsers  get all data...*/
func GetAllUers(p User) ([]User, error) {

	sel := qb.Select("users").Where(qb.Eq("email")).Limit(100).AllowFiltering()

	if p.Password != "" {
		sel.Where(qb.Eq("password"))
	}
	stmt, names := sel.ToCql()

	//	fmt.Println("stmt =", stmt)
	//	fmt.Println("names =", names)

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&p)
	//	fmt.Println("q.Query=  ", q.Query)
	defer q.Release()

	var people []User
	var err error
	if err = gocqlx.Select(&people, q.Query); err != nil {
		fmt.Println("select Err:", err)
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
	fmt.Println(names)
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

	fmt.Println("stmt =", stmt)
	fmt.Println("h =", h)
	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)
	fmt.Println("q =", q)

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
	if h.Hospitalbed != 0 {
		sel.Where(qb.Eq("hospitalbed"))
	}
	if len(h.Patientname) != 0 {
		sel.Where(qb.Eq("patientname"))
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

func GetDeviceInfo(d DeviceInfo) ([]DeviceInfo, error) {
	// Insert with query parameters bound from struct.

	log := logs.GetBeeLogger()
	log.Info("Get Device information")
	fmt.Println("d =", d)

	sel := qb.Select("device_info").Limit(100).AllowFiltering()
	if len(d.Hospitalname) != 0 {
		sel = qb.Select("device_info").Where(qb.Eq("hospitalname")).Limit(100).AllowFiltering()
		if d.Hospitaldeviceid != 0 {
			sel.Where(qb.Eq("Hospitaldeviceid"))
		}
		if len(d.Channelid) != 0 {
			sel.Where(qb.Eq("Channelid"))
		}
		if len(d.Deviceid) != 0 {
			sel.Where(qb.Eq("Deviceid"))
		}
	}

	stmt, names := sel.ToCql()

	fmt.Println("stmt =", stmt)
	fmt.Println("names =", names)

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&d)
	fmt.Println("q.Query=  ", q.Query)
	defer q.Release()

	var deviceinfo []DeviceInfo
	if err := gocqlx.Select(&deviceinfo, q.Query); err != nil {
		fmt.Println("select Err:", err)
	}
	return deviceinfo, nil
}
