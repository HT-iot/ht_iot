package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/gocql/gocql"

	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
)

var tables = []string{
	`CREATE TABLE IF NOT EXISTS hospital_patient_info (
		hospital_name text,
		hospital_id text,
		hospital_zone text,  
		hopsital_bed int,
		patient_name text,
		patient_sex text,
		patient_id text,
		patient_entr_time timeuuid,
		patient_exit_time timeuuid,
		patient_in_hospital boolean,
		device_id text,
		device_hospital_id int,
		channel_id text,
		meta map<text,text>,
		PRIMARY KEY (hospital_name,hospital_zone,patient_name,device_hospital_id,patient_in_hospital)
		)`,

	`CREATE TABLE IF NOT EXISTS device_info (
		manufacturer text,
		model_number text,
		device_id text,
		firmware_version text,
		reboot boolean,
		factory_reset boolean,
		available_power_source int,
		power_source_voltage int,
		power_source_current int,
		battery_level int,
		memory_free int,
		error_code int,
		reset_error_code text,
		current_time timeuuid,
		utc_offset text,
		timezone text,
		supported_binding_modes text,
		device_type text,
		hardware_version text,
		software_version text,
		hospital_name text,
		device_hospital_id int,
		channel_id text,
		PRIMARY KEY (device_id,channel_id,hospital_name,device_hospital_id)
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
	HospitalName     string
	HospitalID       string
	HospitalZone     string
	HopsitalBed      int
	PatientName      string
	PatientSex       string
	PatientID        string
	AdmissionTime    gocql.UUID
	DischargeTime    gocql.UUID
	InHospital       bool
	DeviceID         string
	HospitalDeviceID int
	ChannelID        string
	Meta             map[string]string
}

/*DeviceInfo : The Device information */
type DeviceInfo struct {
	Manufacturer          string
	ModelNumber           string
	DeviceID              string
	FirmwareVersion       string
	Reboot                bool
	FactoryReset          bool
	AvailablePowerSource  int
	PowerSourceVoltage    int
	PowerSourceCurrent    int
	BatteryLevel          int
	MemoryFree            int
	ErrorCode             int
	ResetErrorCode        string
	CurrentTime           gocql.UUID
	UtcOffset             string
	Timezone              string
	SupportedBindingModes string
	DeviceType            string
	HardwareVersion       string
	SoftwareVersion       string
	HospitalName          string
	HospitalDeviceID      int
	ChannelID             string
}

/*GetUers  get data... */
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

/*GetUers  get data... */
/*
func GetUers() (User, error) {
	log := logs.GetBeeLogger()

	cql := `SELECT email, password FROM users LIMIT 1`

	u := User{}
	var err error
	//var email string
	//var password string

	if SessionMgr == nil || SessionMgr.Closed() {
		log.Critical("SessionMgr is nil")
		//return , nil
	}

	if err := SessionMgr.Query(cql).
		Scan(&u.Email, &u.Password); err != nil {

		log.Error("Manager DB query failure")
		//return nil
	}
	log.Info("user talbe, email is " + u.Email + " password is " + u.Password)
	//	return u, err

	/////

	cql = `SELECT channel,id FROM messages_by_channel LIMIT 1`
	//u := User{}
	//var err error
	var channel, id string
	//var password string
	if SessionMsg == nil || SessionMsg.Closed() {
		log.Critical("SessionMgr is nil")
		//return , nil
	}
	if err := SessionMsg.Query(cql).
		Scan(&channel, &id); err != nil {

		log.Error("Manager DB query failure", err.Error())
		//return nil
	}
	log.Info("Message channel, id is " + channel + " password is " + id)
	return u, err

}
*/
/*
func get_clients_by_uer{
	session := createSession(t)
	defer session.Close()

	stmt, names := qb.Select("manager.clients_by_user").ToCql()

	q := gocqlx.Query(session.Query(stmt), names).BindMap(qb.M{
		"first_name": []string{"Patricia", "Igy", "Ian"},
	})
	defer q.Release()

	var people []Person
	if err := gocqlx.Select(&people, q.Query); err != nil {
		t.Fatal("select:", err)
	}

	t.Log(people)
}
*/
