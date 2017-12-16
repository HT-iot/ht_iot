package controllers

import (
	"ht_iot/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/gocql/gocql"
)

var Hospitalslice = make([]models.HospitalPatientInfo, 0, 10)

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

	var h models.HospitalPatientInfo
	op := this.Input().Get("op")
	id, _ := this.GetInt("id")

	switch op {
	case "add":
		{
			h.Hospitalname = this.GetString("hospital_name")
			h.Hospitalzone = this.GetString("hospital_zone")
			h.Hospitalbed, _ = this.GetInt("hospital_bed")
			h.Patientname = this.GetString("p_name")
			h.Patientsex = this.GetString("p_sex")
			h.Hospitaldeviceid, _ = this.GetInt("p_device")
			h.Inhospital = true
			h.Patiententrtime = time.Now().Local()
			Hospitalslice = append(Hospitalslice, h)

			break
		}
	case "del":
		{
			Hospitalslice = append(Hospitalslice[:id], Hospitalslice[id+1:]...)
			break
		}
	case "wrte":
		{
			Hospitalslice[id].Id = gocql.TimeUUID()
			h = Hospitalslice[id]
			_ = models.InsertPatient(h)
			Hospitalslice = append(Hospitalslice[:id], Hospitalslice[id+1:]...)
			break

		}

	}

	this.Data["Hospitalsilce"] = &Hospitalslice
	this.TplName = "pconfig.html"

}

func (this *PconfigController) Post() {
	this.TplName = "pconfig.html"

}
