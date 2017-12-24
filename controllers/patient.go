package controllers

import (
	"ht_iot/models"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type PatientController struct {
	beego.Controller
}

func (this *PatientController) Get() {
	this.TplName = "patient_status.html"
	this.Data["IsSconfig"] = true

	flag := checkAccount(this.Ctx)
	this.Data["ISLogin"] = flag
	if !flag {
		this.Redirect("/login", 302)
		return
	}

	var hPatients []models.HospitalPatientInfo

	name := this.GetString("hospital_name")
	zone := this.GetString("hospital_zone")
	pName := this.GetString("p_name")

	if name != "" || zone != "" || pName != "" {
		hp := models.HospitalPatientInfo{
			//Hospitalname: name,
			//Hospitalzone: zone,
			Patientname: pName}
		hPatients, _ = models.GetPatient(hp)
	} else {
		hPatients, _ = models.GetAllPatient()
	}

	if len(hPatients) == 0 {
		logs.Error("Can't get patient info from DB")
	}

	Patients := make([]models.PatientInfo, len(hPatients))
	for i, p := range hPatients {
		Patients[i].Hospitalname = p.Hospitalname
		Patients[i].Hospitalzone = p.Hospitalzone
		Patients[i].Hospitalbed = p.Hospitalbed
		Patients[i].Patientname = p.Patientname
		Patients[i].Hospitaldeviceid = p.Hospitaldeviceid
		Patients[i].Channelid = p.Channelid
		Patients[i].Deviceid = p.Deviceid

		err := models.GetMsg(&Patients[i])
		if err != nil {
			logs.Error("Get the status failure!")
		}
		logs.Debug(Patients[i])
	}

	this.Data["Patients"] = &Patients

	this.TplName = "patient_status.html"

}

func (this *PatientController) Post() {
	this.TplName = "patient_status.html"
}
