package controllers

import (
	"fmt"
	"ht_iot/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type DeviceinfoController struct {
	beego.Controller
}

var id = 0
var d models.DeviceInfo
var Deviceinfo []models.DeviceInfo
var err error

func (this *DeviceinfoController) Get() {
	this.TplName = "deviceinfo.html"
	this.Data["IsDconfig"] = true

	flag := checkAccount(this.Ctx)
	this.Data["ISLogin"] = flag
	if !flag {
		this.Redirect("/login", 302)
		return
	}
}

/*
	op := this.Input().Get("op")

	fmt.Println("op", op)

	switch op {
	case "Check":
		{
			//		d = this.getfromhtml()
			d.Hospitalname = this.GetString("hospital_name")

			fmt.Println("Hospitalname", d.Hospitalname)
			this.Data["Hospitalname"] = d.Hospitalname
			d.Hospitaldeviceid = this.GetString("hospitaldeviceid")
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
*/

func (this *DeviceinfoController) Post() {
	this.TplName = "deviceinfo.html"
	logs.Debug("Device information to Device table")
	var Mystruct models.DeviceTable
	p := models.DeviceInfo{}
	Deviceinfo, err = models.GetDeviceInfo(p)
	if len(Deviceinfo) > 0 {
		Mystruct.Data = append(Mystruct.Data, Deviceinfo...)
	} else {
		Mystruct.Data = nil
	}
	this.Data["json"] = &Mystruct
	this.ServeJSON()
}

func (this *DeviceinfoController) GetModal() {

	logs.Debug("Device information to Device Modal")
	deviceid := this.GetString("deviceid")
	var Mystruct models.DeviceTable
	var devicemodal []models.DeviceInfo
	p := models.DeviceInfo{Deviceid: deviceid}

	devicemodal, err = models.GetDeviceInfo(p)
	fmt.Println(devicemodal)
	if len(devicemodal) == 1 {
		Mystruct.Data = append(Mystruct.Data, devicemodal...)
	} else {
		Mystruct.Data = nil
	}
	this.Data["json"] = &devicemodal
	this.ServeJSON()
}

func (this *DeviceinfoController) PostModal() {
	logs.Debug("Get Device information from Device Modal")
	//	deviceid := this.GetString("deviceid")
	var Getstruct In
	d = this.getfromhtml()
	//	hd := this.GetString("deviceid")
	//	d.Id = uuid.FromString(d.Deviceid)
	fmt.Println("PostModal d=", d)
	err := models.UpdateDeviceItem(d)

	if err == nil {
		Getstruct.Info = "更新成功"
		Getstruct.Succ = "Succ"
	} else {
		Getstruct.Info = "更新失败"
		Getstruct.Succ = "Fail"
	}
	this.Data["json"] = &Getstruct
	this.ServeJSON()
}

func (this *DeviceinfoController) PostMerge() {
	logs.Debug("Device information to Device table")
	str0 := this.GetString("hospitalname")

	var Getstruct In
	//	fmt.Println("str", str0, str1, str2)
	_ = models.InputDevices(str0)
	p := models.DeviceInfo{Hospitalname: d.Hospitalname}
	Deviceinfo, _ = models.GetDeviceInfo(p)
	fmt.Println("Deviceinfo", Deviceinfo)
	if len(Deviceinfo) > 0 {
		//				d = Deviceinfo[id]
		Getstruct.Info = "合并终端成功"
		Getstruct.Succ = "Succ"
	} else {
		Getstruct.Info = "无终端可添加"
		Getstruct.Succ = "Fail"
	}
	this.Data["json"] = &Getstruct
	this.ServeJSON()
}

func (this *DeviceinfoController) getfromhtml() models.DeviceInfo {
	var h models.DeviceInfo
	h.Hospitalname = this.GetString("hospitalname")
	h.Hospitaldeviceid = this.GetString("hospitaldeviceid")
	h.Deviceid = this.GetString("deviceid")
	h.Channelid = this.GetString("channelid")

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
