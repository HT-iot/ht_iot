package controllers

import (
	"fmt"
	"ht_iot/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gocql/gocql"
)

var Hospitalslice []models.HospitalPatientInfo

type In struct {
	Succ string `json:"succ"`
	Info string `json:"info"`
}

/* PconfigController...*/
type PconfigController struct {
	beego.Controller
}

/* Get...*/
func (this *PconfigController) Get() {
	this.TplName = "pconfig.html"
	this.Data["IsPconfig"] = true

	flag := checkAccount(this.Ctx)
	this.Data["ISLogin"] = flag
	if !flag {
		this.Redirect("/login", 302)
		return
	}
	//	fmt.Println("Hospital =", Hospital)
}

func (this *PconfigController) Post() {
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

func (this *PconfigController) PostPat() {
	logs.Debug("Input the data in Pconfig")
	var hv bool
	var Getstruct In
	var h models.HospitalPatientInfo
	{
		h.Hospitalname = this.GetString("hospitalname")
		h.Hospitalzone = this.GetString("hospitalzone")
		h.Hospitalbed = this.GetString("hospitalbed")
		h.Patientname = this.GetString("patientname")
		h.Patientsex = this.GetString("patientsex")
		h.Patientid = this.GetString("patientid")
		h.Hospitaldeviceid = this.GetString("hospitaldeviceid")
	}

	h.Id = gocql.TimeUUID()
	h.Patiententrtime = time.Now()

	h.Channelid, h.Deviceid, hv = models.GetPaientIDs(h.Hospitalname, h.Hospitaldeviceid)
	if hv {
		_ = models.InsertPatient(h)
		Getstruct.Info = "添加成功"
		Getstruct.Succ = "succ"
	} else {
		Getstruct.Info = "无该医院终端ID, 添加失败"
		Getstruct.Succ = "fail"
	}

	this.Data["json"] = &Getstruct
	this.ServeJSON()
}

func (this *PconfigController) GetPat() {
	logs.Debug("the Patient information")
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

func (this *PconfigController) PostLine() {
	logs.Debug("Update Line")

	var hv bool
	var Getstruct In
	var h models.HospitalPatientInfo
	{
		h.Hospitalname = this.GetString("hospitalname")
		h.Hospitalzone = this.GetString("hospitalzone")
		h.Hospitalbed = this.GetString("hospitalbed")
		h.Patientname = this.GetString("patientname")
		h.Patientsex = this.GetString("patientsex")
		h.Patientid = this.GetString("patientid")
		h.Hospitaldeviceid = this.GetString("hospitaldeviceid")
		h.Deviceid = this.GetString("deviceid")
		h.Channelid = this.GetString("channelid")
	}

	//	h.Id = gocql.TimeUUID()
	//	h.Patiententrtime = time.Now()
	fmt.Println("h=", h)

	hv = models.UpdatePatient(h)
	if hv {
		//	_ = models.InsertPatient(h)
		Getstruct.Info = "添加成功"
		Getstruct.Succ = "succ"
	} else {
		Getstruct.Info = "无该医院终端ID, 添加失败"
		Getstruct.Succ = "fail"
	}

	this.Data["json"] = &Getstruct
	this.ServeJSON()

}

func (this *PconfigController) GetLine() {
	logs.Debug("Post Line")

	//	var d models.HospitalPatientInfo
	var err error
	var Mystruct models.HospitalTable
	//	d.Hospitaldeviceid = Hospitalslice[id].Hospitaldeviceid

	//		h.Patiententrtime = time.Now()
	//	d.Hospitalname = this.GetString("hospitalname")
	device := this.GetString("deviceid")
	d := models.HospitalPatientInfo{
		//Hospitalname: name,
		//Hospitalzone: zone,
		Deviceid: device}

	fmt.Println("PostLine d Patientid=", d)

	Hospitalslice, err = models.GetPatient(d)

	if err == nil {
		Mystruct.Data = append(Mystruct.Data, Hospitalslice...)
	} else {
		Mystruct.Data = nil
	}
	//	fmt.Println("PostLine=", Mystruct)
	this.Data["json"] = &Hospitalslice
	this.ServeJSON()
	/*
			if (err == nil) {
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
	*/
}
