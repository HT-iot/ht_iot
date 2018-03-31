package controllers
//该代码缺一致性检查
import (
//	"fmt"
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


func (this *DeviceinfoController) Post() {
	logs.Debug("Device information to Device table")
	this.TplName = "deviceinfo.html"

	var Mystruct models.DeviceTable
	p := models.DeviceInfo{}
	Deviceinfo, err = models.GetDeviceInfo(p)

	Mystruct.Data = make([]models.DeviceInfo, len(Deviceinfo))
	copy(Mystruct.Data,Deviceinfo)
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

	if len(devicemodal) == 1 {
		Mystruct.Data = append(Mystruct.Data, devicemodal...)
	} else {
		Mystruct.Data = nil
	}
	this.Data["json"] = &devicemodal
	this.ServeJSON()
}

func (this *DeviceinfoController) PostModal() {
	logs.Debug("POst Updated Device information to DB")

	var Getstruct In
	d = this.getfromhtml()

//	fmt.Println("PostModal d=", d)
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
	logs.Debug("Got the created Device to Device table from DB-channels_by_user")
	str0 := this.GetString("hospitalname")

	var Getstruct In
	_ = models.InputDevices(str0)
	p := models.DeviceInfo{Hospitalname: d.Hospitalname}
	Deviceinfo, _ = models.GetDeviceInfo(p)
//	fmt.Println("Deviceinfo", Deviceinfo)
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

