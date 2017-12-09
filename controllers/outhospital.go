package controllers

import (
	"fmt"
	"ht_iot/models"

	"github.com/astaxie/beego"
)

type OuthospitalController struct {
	beego.Controller
}

func (this *OuthospitalController) Get() {
	this.TplName = "outhospital.html"
	this.Data["IsOconfig"] = true

	flag := checkAccount(this.Ctx)
	this.Data["ISLogin"] = flag
	if !flag {
		this.Redirect("/login", 302)
		return
	}

	/*
		fmt.Println("IsLogin0:  ", IsLogin)
		this.Data["IsOconfig"] = true
		this.TplName = "outhospital.html"
		//	this.Data["ISLogin"] = IsLogin
		flag := checkAccount(this.Ctx)
		fmt.Println("O flag:  ", flag)
		this.Data["ISLogin"] = flag
		if !flag {
			this.Redirect("/login", 302)
			return
		}

			if !IsLogin {
				this.Redirect("/login", 302)
				return
			}
	*/
	/*	flag := checkAccount(this.Ctx)
		this.Data["IsLogin"] = flag
		if !flag {
			this.Redirect("/login", 302)
			return
		}
	*/
	var h models.HospitalPatientInfo
	var Hospitalslicetmp []models.HospitalPatientInfo

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
			*/
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
			*/break

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

}

func (this *OuthospitalController) Post() {
	this.TplName = "outhospital.html"
}
