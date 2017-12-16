package controllers

import (
	"fmt"
	"ht_iot/models"

	"github.com/astaxie/beego"
)

type DeviceinfoController struct {
	beego.Controller
}

func (this *DeviceinfoController) Get() {
	this.TplName = "deviceinfo.html"
	this.Data["IsDconfig"] = true

	flag := checkAccount(this.Ctx)
	this.Data["ISLogin"] = flag
	if !flag {
		this.Redirect("/login", 302)
		return
	}

	var d models.DeviceInfo
	var Deviceinfo []models.DeviceInfo

	d.Hospitalname = this.GetString("hospital_name")
	d.Hospitaldeviceid, _ = this.GetInt("Hospitaldeviceid")
	d.Channelid = this.GetString("Channelid")
	d.Deviceid = this.GetString("Deviceid")

	p := models.DeviceInfo{Hospitalname: d.Hospitalname, Hospitaldeviceid: d.Hospitaldeviceid, Channelid: d.Channelid, Deviceid: d.Deviceid}
	Deviceinfo, _ = models.GetDeviceInfo(p)
	fmt.Println(Deviceinfo)
	/*
		op := this.Input().Get("op")
		id, _ := this.GetInt("id")

		switch op {
		case "check":
			{
				h.Hospitalname = this.GetString("hospital_name")
				h.Hospitalzone = this.GetString("hospital_zone")
				h.Hospitalbed, _ = this.GetInt("hospital_bed")
				h.Patientname = this.GetString("p_name")
				h.Patientsex = this.GetString("p_sex")
				h.Hospitaldeviceid, _ = this.GetInt("p_device")
				/*			p := models.User{Email: "yong@test.com"}
							q, _ := models.GetAllUers(p)
							fmt.Println(q)

				fmt.Println("h=", h.Hospitalname)
				p := models.HospitalPatientInfo{Hospitalname: h.Hospitalname, Hospitalzone: h.Hospitalzone, Patientname: h.Patientname}
				Hospitalslicetmp, _ = models.GetPatient(p)
				fmt.Println("h=", Hospitalslicetmp)
				/*
					for _, d := range Hospitalslice {
						if ((len(h.HospitalID) == 0) || (h.HospitalID == d.HospitalID)) &&
							((len(h.HospitalZone) == 0) || (h.HospitalZone == d.HospitalZone)) &&
							(((h.HospitalBed) == 0) || (h.HospitalBed == d.HospitalBed)) &&
							((len(h.PatientName) == 0) || (h.PatientName == d.PatientName)) &&
							(((h.HospitalDeviceID) == 0) || (h.HospitalDeviceID == d.HospitalDeviceID)) {
							Hospitalslicetmp = append(Hospitalslicetmp, d)
						}

					}
				break

			}

		case "del":
			{
				Hospitalslice = append(Hospitalslice[:id], Hospitalslice[id+1:]...)
				break
			}
		}
		//	this.Ctx.WriteString(fmt.Sprint(Hospitalslice))
		//	fmt.Println(Hospitalslice)
		//	this.Data["IsPconfig"] = true
		this.Data["Hospitalsilce"] = &Hospitalslicetmp
		fmt.Println("IsLogin1:  ", IsLogin)
		this.TplName = "outhospital.html"
	*/
}

func (this *DeviceinfoController) Post() {
	this.TplName = "deviceinfo.html"
}
