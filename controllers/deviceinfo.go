package controllers

import (
	"fmt"
	"ht_iot/models"
	"time"

	"github.com/astaxie/beego"
)

type DeviceinfoController struct {
	beego.Controller
}

var id = 0
var d models.DeviceInfo
var Deviceinfo []models.DeviceInfo
var err error

func (this *DeviceinfoController) Prepare() {
	this.Data["IsDconfig"] = true

	flag := checkAccount(this.Ctx)
	this.Data["ISLogin"] = flag
	if !flag {
		this.Redirect("/login", 302)
		return
	}
	//	p := models.DeviceInfo{Hospitalname: "", Hospitaldeviceid: 0, Channelid: "", Deviceid: ""}
	p := models.DeviceInfo{}
	Deviceinfo, err = models.GetDeviceInfo(p)

	if len(Deviceinfo) <= 0 {
		id = -1
		fmt.Println("err", err)
	}

}

func (this *DeviceinfoController) Get() {

	op := this.Input().Get("op")

	fmt.Println("op", op)

	switch op {
	case "Check":
		{
			//		d = this.getfromhtml()
			d.Hospitalname = this.GetString("hospital_name")

			fmt.Println("Hospitalname", d.Hospitalname)
			this.Data["Hospitalname"] = d.Hospitalname
			d.Hospitaldeviceid, err = this.GetInt("hospitaldeviceid")
			d.Channelid = this.GetString("channel_id")
			d.Deviceid = this.GetString("device_id")
			p := models.DeviceInfo{Hospitalname: d.Hospitalname, Hospitaldeviceid: d.Hospitaldeviceid,
				Channelid: d.Channelid, Deviceid: d.Deviceid}
			fmt.Println("p", p)
			Deviceinfo, _ = models.GetDeviceInfo(p)
			if len(Deviceinfo) > 0 {
				id = 0
				//				d = Deviceinfo[id]
				this.writetohtml(id)
			} else {
				id = -1
				this.Data["hospital_name_v"] = d.Hospitalname
				this.Data["hospitaldeviceid_v"] = d.Hospitaldeviceid
				this.Data["device_id_v"] = d.Deviceid
				this.Data["channel_id_v"] = d.Channelid
			}
			break
		}

	case "Update":
		{
			d = this.getfromhtml()
			d.Id = Deviceinfo[id].Id
			_ = models.UpdateDeviceItem(d)
			//	this.Data["hospital_value"] = Deviceinfo[id].Hospitalname
			Deviceinfo, _ = models.GetDeviceInfo(d)
			if len(Deviceinfo) > 0 {
				id = 0
				//				d = Deviceinfo[id]
				this.writetohtml(id)
			} else {
				id = -1
				this.Data["hospital_name_v"] = d.Hospitalname
				this.Data["hospitaldeviceid_v"] = d.Hospitaldeviceid
				this.Data["device_id_v"] = d.Deviceid
				this.Data["channel_id_v"] = d.Channelid
			}
			break
		}
	case "Edit":
		{
			id, _ = this.GetInt("id")
			fmt.Println("id", id)
			this.writetohtml(id)
			break
		}
	case "Merge":
		{
			str := this.Input().Get("hospital_name")
			fmt.Println("str", str)
			_ = models.InputDevices(str)
			p := models.DeviceInfo{Hospitalname: d.Hospitalname}
			Deviceinfo, _ = models.GetDeviceInfo(p)
			fmt.Println("Deviceinfo", Deviceinfo)
			if len(Deviceinfo) > 0 {
				id = 0
				//				d = Deviceinfo[id]
				this.writetohtml(id)
			} else {
				id = -1
			}
			break

			break
		}
	default:
		{
			if id >= 0 {
				this.writetohtml(id)
			}
			break

		}
	}
	//	this.Ctx.WriteString(fmt.Sprint(Hospitalslice))
	//	fmt.Println(Hospitalslice)
	//	this.Data["IsPconfig"] = true
	//	this.Data["Deviceinfo"] = &Deviceinfo
	//		fmt.Println("IsLogin1:  ", IsLogin)
	this.Data["Deviceinfo"] = &Deviceinfo
	this.TplName = "deviceinfo.html"
}

func (this *DeviceinfoController) Post() {
	this.TplName = "deviceinfo.html"
}

func (this *DeviceinfoController) writetohtml(id int) {
	d.Hospitalname = this.GetString("hospital_name")
	d.Hospitaldeviceid, _ = this.GetInt("hospitaldeviceid")
	fmt.Println(Deviceinfo[id])

	this.Data["hospital_name_v"] = Deviceinfo[id].Hospitalname
	this.Data["hospitaldeviceid_v"] = Deviceinfo[id].Hospitaldeviceid
	this.Data["device_id_v"] = Deviceinfo[id].Deviceid
	this.Data["channel_id_v"] = Deviceinfo[id].Channelid

	this.Data["manufacturer_v"] = Deviceinfo[id].Manufacturer
	this.Data["modelnumber_v"] = Deviceinfo[id].Modelnumber
	this.Data["firmwareversion_v"] = Deviceinfo[id].Firmwareversion
	this.Data["devicetype_v"] = Deviceinfo[id].Devicetype

	this.Data["availablepowersource_v"] = Deviceinfo[id].Availablepowersource
	this.Data["powersourcevoltage_v"] = Deviceinfo[id].Powersourcevoltage
	this.Data["powersourcesurrent_v"] = Deviceinfo[id].Powersourcesurrent
	this.Data["batterylevel_v"] = Deviceinfo[id].Batterylevel

	this.Data["supportedbindingmodes_v"] = Deviceinfo[id].Supportedbindingmodes
	this.Data["hardwareversion_v"] = Deviceinfo[id].Hardwareversion
	this.Data["softwareversion_v"] = Deviceinfo[id].Softwareversion
}

func (this *DeviceinfoController) getfromhtml() models.DeviceInfo {
	var h models.DeviceInfo
	h.Hospitalname = this.GetString("hospital_name")
	h.Hospitaldeviceid, _ = this.GetInt("hospitaldeviceid")
	h.Deviceid = this.GetString("device_id")
	h.Channelid = this.GetString("channel_id")

	h.Manufacturer = this.GetString("manufacturer")
	h.Modelnumber = this.GetString("modelnumber")
	h.Firmwareversion = this.GetString("firmwareversion")
	h.Devicetype = this.GetString("devicetype")

	h.Availablepowersource, _ = this.GetInt("availablepowersource")
	h.Powersourcevoltage, _ = this.GetInt("powersourcevoltage")
	h.Powersourcesurrent, _ = this.GetInt("powersourcesurrent")
	h.Batterylevel, _ = this.GetInt("batterylevel")

	h.Supportedbindingmodes = this.GetString("supportedbindingmodes")
	h.Hardwareversion = this.GetString("hardwareversion")
	h.Softwareversion = this.GetString("softwareversion")
	h.Currenttime = time.Now()

	return h
}
