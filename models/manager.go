package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
)

var tables = []string{
	`CREATE TABLE IF NOT EXISTS manager.hospital_by_patient (
		hospitalname text,
		hospitalzone text,
		patientname text,
		hospitaldeviceid int,
		inhospital boolean,
		channelid text,
		deviceid text,
		hospitalbed int,
		hospitalid text,
		patiententrtime timeuuid,
		patientexittime timeuuid,
		patientid text,
		patientsex text,
		meta map<text, text>,
		PRIMARY KEY (hospitalname, hospitalzone, patientname, hospitaldeviceid, inhospital)
		)`,

	`CREATE TABLE IF NOT EXISTS manager.device_info (
		Manufacturer  text,
		Modelnumber text,
		Deviceid text,
		FirmwareVersion text,
		reboot boolean,
		factoryreset boolean,
		Availablepowersource int,
		Powersourcevoltage int,
		Powersourcesurrent int,
		Batterylevel int,
		Memoryfree int,
		Errorcode int,
		Reseterrorcode text,
		Currenttime timeuuid,
		Utcoffset text,
		Timezone text,
		Supportedbindingmodes text,
		Devicetype text,
		Hardwareversion text,
		Softwareversion text,
		Hospitalname text,
		Hospitaldeviceid int,
		Channelid text,
		PRIMARY KEY (Deviceid,Channelid,Hospitalname,Hospitaldeviceid)
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
	Hospitalname     string
	Hospitalid       string
	Hospitalzone     string
	Hospitalbed      int
	Patientname      string
	Patientsex       string
	Patientid        string
	Patiententrtime  string
	Patientexittime  string
	Inhospital       bool
	Deviceid         string
	Hospitaldeviceid int
	Channelid        string
	Meta             map[string]string
}

/*DeviceInfo : The Device information */
type DeviceInfo struct {
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
	ReseterrorCode        string
	Currenttime           gocql.UUID
	UtcOffset             string
	Timezone              string
	Supportedbindingmodes string
	Devicetype            string
	Hardwareversion       string
	Softwareversion       string
	Hospitalname          string
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

	fmt.Println("stmt =", stmt)
	fmt.Println("names =", names)

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&p)
	fmt.Println("q.Query=  ", q.Query)
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

	stmt, names := qb.Insert("hospital_by_patient").Columns("hospitalname", "hospitalzone",
		"patientname", "hospitaldeviceid", "inhospital", "channelid",
		"deviceid", "hospitalbed", "hospitalid", "patiententrtime",
		"patientid", "patientsex", "meta").
		ToCql()

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)

	if err := q.ExecRelease(); err != nil {
		log.Critical("select:" + err.Error())
		return err
	}
	return nil
}

func GetPatient(h HospitalPatientInfo) ([]HospitalPatientInfo, error) {
	// Insert with query parameters bound from struct.

	log := logs.GetBeeLogger()
	log.Info("Get Patient information")

	sel := qb.Select("hospital_by_patient").Where(qb.Eq("hospitalname")).Limit(100).AllowFiltering()
	/*
		if len(h.Hospitalzone) != 0 {
			sel.Where(qb.Eq("hospitalzone"))
		}
	*/
	if len(h.Patientname) != 0 {
		sel.Where(qb.Eq("patientname"))
	}

	stmt, names := sel.ToCql()

	fmt.Println("stmt =", stmt)
	fmt.Println("h =", h)

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&h)
	fmt.Println("q.Query=  ", q.Query)
	defer q.Release()

	var patient []HospitalPatientInfo
	if err := gocqlx.Select(&patient, q.Query); err != nil {
		fmt.Println("select Err:", err)
	}
	return patient, nil
}
