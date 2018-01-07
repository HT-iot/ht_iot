package controllers

import (
	"fmt"
	"ht_iot/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gocql/gocql"
)

//var Hospitalslice []models.HospitalPatientInfo

/* OuthospitalController...*/
type OuthospitalController struct {
	beego.Controller
}

/* Get...*/
func (this *OuthospitalController) Get() {
	this.TplName = "outhospital.html"
	this.Data["Isouthospital"] = true

	flag := checkAccount(this.Ctx)
	this.Data["ISLogin"] = flag
	if !flag {
		this.Redirect("/login", 302)
		return
	}
	//	fmt.Println("Hospital =", Hospital)
}

func (this *OuthospitalController) GetPat() {
	logs.Debug("Input the data in Pconfig")

	var h models.HospitalPatientInfo
	{
		h.Hospitalname = this.GetString("hospitalname")
		h.Hospitalzone = this.GetString("hospitalzone")
		h.Hospitalbed = this.GetString("hospitalbed")
		h.Patientname = this.GetString("patientname")
		h.Hospitaldeviceid = this.GetString("hospitaldeviceid")
	}
	Hospitalslice = append(Hospitalslice, h)
	Getstruct := In{Succ: "add", Info: ""}
	this.Data["json"] = &Getstruct
	this.ServeJSON()
}

func (this *OuthospitalController) PostPat() {
	logs.Debug("Read the Patient information")
	var Mystruct models.HospitalTable
	var err error
	Hospitalslice, err = models.GetAllPatient()

	if err == nil {
		Mystruct.Data = append(Mystruct.Data, Hospitalslice...)
	} else {
		Mystruct.Data = nil
	}

	fmt.Println(Mystruct)
	this.Data["json"] = &Mystruct
	this.ServeJSON()
}

func (this *OuthospitalController) Line() {
	logs.Debug("Line change in Pconfig table")
	var d models.HospitalPatientInfo
	var Getstruct In
	var hv bool

	actions := this.GetString("action")
	id, _ := this.GetInt("id")

	if actions == "Add" {

		d.Id = gocql.TimeUUID()
		d.Patiententrtime = time.Now()
		d.Hospitalname = Hospitalslice[id].Hospitalname
		d.Hospitalzone = Hospitalslice[id].Hospitalzone
		d.Hospitalbed = Hospitalslice[id].Hospitalbed
		d.Patientname = Hospitalslice[id].Patientname
		d.Hospitaldeviceid = Hospitalslice[id].Hospitaldeviceid

		//		h.Patiententrtime = time.Now()
		d.Channelid, d.Deviceid, hv = models.GetPaientIDs(d.Hospitalname, d.Hospitaldeviceid)
		if hv {
			_ = models.InsertPatient(d)
			Getstruct.Info = "添加成功"
			Getstruct.Succ = "add"
			if len(Hospitalslice) > 1 {
				Hospitalslice = append(Hospitalslice[:id], Hospitalslice[id+1:]...)
				Getstruct.Succ = "add"
				Getstruct.Info = "注册成功"
			} else {
				Hospitalslice = nil
				Getstruct.Succ = "nil"
				Getstruct.Info = "注册成功"
			}
		} else {
			Getstruct.Info = "无该医院终端ID, 添加失败"
			Getstruct.Succ = "add"
		}
		//	mystruct = Out{Succ: str, Refresh: Out.Refresh}
	}
	if actions == "Del" {
		if len(Hospitalslice) > 1 {
			Hospitalslice = append(Hospitalslice[:id], Hospitalslice[id+1:]...)
			Getstruct.Info = "取消成功"
			Getstruct.Succ = "del"
		} else {
			Hospitalslice = nil
			Getstruct.Info = "取消成功"
			Getstruct.Succ = "nil"
		}

	}
	this.Data["json"] = &Getstruct
	this.ServeJSON()
}
