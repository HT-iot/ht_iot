package controllers

import (
	"ht_iot/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type StatusController struct {
	beego.Controller
}

/*Get reponse the status query */
func (s *StatusController) Get() {
	s.Data["IsSconfig"] = true

	flag := checkAccount(s.Ctx)
	s.Data["ISLogin"] = flag
	if !flag {
		s.Redirect("/login", 302)
		return
	}

	s.TplName = "patient_status.html"
}

/*GetStatus reponse the API query */
func (s *StatusController) GetStatus() {
	logs.Debug("GetStatus")
	var hPatients []models.HospitalPatientInfo

	name := s.GetString("hospital_name")
	zone := s.GetString("hospital_zone")
	pName := s.GetString("p_name")

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

	logs.Debug(Patients)

	var data models.DataTable
	data.Data = Patients

	s.Data["json"] = &data
	s.ServeJSON()
}

/*PostStatus reponse the API query */
func (s *StatusController) PostStatus() {

}
