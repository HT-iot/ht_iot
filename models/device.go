package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
)

func GetDeviceInfo(d DeviceInfo) ([]DeviceInfo, error) {
	// Insert with query parameters bound from struct.

	log := logs.GetBeeLogger()
	log.Info("Get Device information")

	sel := qb.Select("device_info").AllowFiltering()
	if len(d.Hospitalname) > 0 {
		sel = qb.Select("device_info").Where(qb.Eq("hospitalname")).AllowFiltering()
		if len(d.Hospitaldeviceid) > 0 {
			sel.Where(qb.Eq("hospitaldeviceid"))
		}
		if len(d.Channelid) > 0 {
			sel.Where(qb.Eq("channelid"))
		}
		if len(d.Deviceid) > 0 {
			sel.Where(qb.Eq("deviceid"))
		}
	} else {
		if len(d.Deviceid) > 0 {
			sel = qb.Select("device_info").Where(qb.Eq("deviceid")).AllowFiltering()
		}
	}

	//
	stmt, names := sel.ToCql()
	fmt.Println("stmt,names =", stmt, names)
	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&d)
	defer q.Release()

	var deviceinfo []DeviceInfo
	if err := gocqlx.Select(&deviceinfo, q.Query); err != nil {
		fmt.Println("select Err:", err)
		return nil, err
	}
	return deviceinfo, nil
}

func UpdateDeviceItem(d DeviceInfo) error {
	// // Easy update with all parameters bound from struct.

	//		p.Email = append(p.Email, "patricia1.citzen@gocqlx_test.com")
	stmt, names := qb.Update("device_info").
		Set("manufacturer",
			"modelnumber", "firmwareversion", "reboot",
			"factoryreset", "availablepowersource", "powersourcevoltage", "powersourcesurrent",
			"batterylevel", "memoryfree", "errorcode", "reseterrorcode", "currenttime", "utcoffset",
			"timezone", "supportedbindingmodes", "devicetype", "hardwareversion", "softwareversion",
			"hospitaldeviceid").
		Where(qb.Eq("hospitalname"), qb.Eq("channelid"), qb.Eq("deviceid")).
		ToCql()
	fmt.Println("Update d=", d)

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&d)

	fmt.Println("q=", q)

	if err := q.ExecRelease(); err != nil {
		fmt.Println("err=", err)
		return err
	}

	return nil
}

func InputDevices(hospital_name string) error {

	p := ChannelsByUser{Name: hospital_name}
	sel := qb.Select("channels_by_user").Where(qb.Eq("name")).AllowFiltering()

	stmt, names := sel.ToCql()

	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&p)
	//	fmt.Println("q.Query=  ", q.Query)
	defer q.Release()

	var people []ChannelsByUser

	var d DeviceInfo
	var err error

	if err = gocqlx.Select(&people, q.Query); err != nil {
		fmt.Println("select Err:", err)
		return err
	}
	fmt.Println("people:", people)
	for _, v := range people {
		for _, w := range v.Connected {
			if CheckDevice(w) != true {
				//				fmt.Println("device:", v)
				d.Hospitalname = v.Name
				//				d.Id, _ = gocql.ParseUUID(w)
				d.Deviceid = w
				d.Channelid = v.ID.String()
				_ = InsertDevice(d)
			}
		}
	}
	return nil
}

func CheckDevice(device string) bool {
	log := logs.GetBeeLogger()
	log.Info("Get Device information")

	d := DeviceInfo{Deviceid: device}

	sel := qb.Select("device_info").Where(qb.Eq("deviceid")).Limit(100).AllowFiltering()
	//	fmt.Println("d:", d)
	stmt, names := sel.ToCql()
	//	fmt.Println("stmt,names", stmt, names)
	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&d)
	//	fmt.Println("q:", q)
	defer q.Release()
	var deviceinfo []DeviceInfo
	if err := gocqlx.Select(&deviceinfo, q.Query); err != nil {
		fmt.Println("select Err:", err)
	}
	if len(deviceinfo) <= 0 {
		return false
	}
	return true
}

func InsertDevice(d DeviceInfo) bool {
	log := logs.GetBeeLogger()
	log.Info("Insert Device information")

	stmt, names := qb.Insert("device_info").Columns("hospitalname", "manufacturer",
		"modelnumber", "deviceid", "firmwareversion", "reboot",
		"factoryreset", "availablepowersource", "powersourcevoltage", "powersourcesurrent",
		"batterylevel", "memoryfree", "errorcode", "reseterrorcode", "currenttime", "utcoffset",
		"timezone", "supportedbindingmodes", "devicetype", "hardwareversion", "softwareversion",
		"hospitaldeviceid", "channelid",
	).
		ToCql()

	fmt.Println("stmt =", stmt)
	q := gocqlx.Query(SessionMgr.Query(stmt), names).BindStruct(&d)
	fmt.Println("q =", q)

	if err := q.ExecRelease(); err != nil {
		log.Critical("select:" + err.Error())
		return false
	}
	return true
}
