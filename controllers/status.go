package controllers

import (
	"fmt"
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
	fmt.Println("Patients",Patients)
	s.Data["json"] = &data
	s.ServeJSON()
}

/*PostStatus reponse the API query */
func (s *StatusController) PostStatus() {
	logs.Debug("Input the data in Pconfig")
	col := [][2]string{
		{"runstatus","运行状态"},          
		{"hospitalname","医院名称"},
		{"hospitalzone", "病区号"}, 
		{"hospitalbed", "病床号"}, 
		{"patientname", "姓名"},
		{"hospitaldeviceid", "终端号"}, 
		{"puls", "脉搏"}, 
		{"oxgen", "血氧"},
		{"pressurehigh", "收缩压"},
		{"pressurelow", "舒张压"},
		{"longitude", "经度"},
		{"latitude", "纬度"},
	}

	s.Data["json"] = &col
	s.ServeJSON()
}

func (this *StatusController) PostWarnStatus() {
	logs.Debug("Update Line")

	var hv bool
	var Getstruct In
	var h models.Warnpara

	{
		h.Pulsmin,_ = this.GetInt16("pulsmin")
		h.Pulsmax,_ = this.GetInt16("pulsmax")
		h.Oxgenmin,_ = this.GetInt16("oxgenmin")
		h.Oxgenmax,_ = this.GetInt16("oxgenmax")
		h.Pressurelowmin,_ = this.GetInt16("pressurelowmin")
		h.Pressurelowmax,_ = this.GetInt16("pressurelowmax")
		h.Pressurehighmin,_ = this.GetInt16("pressurehighmin")
		h.Pressurehighmax,_ = this.GetInt16("pressurehighmax")
		h.Monitoraddress = this.GetString("monitoraddress")
		h.Monitorradius,_ = this.GetInt16("monitorradius")
//		h.MonitorLongitude,_ = this.GetFloat("monitorlongitude")
//		h.MonitorLatitude,_ = this.GetFloat("monitorlatitude")
	}

	//	h.Id = gocql.TimeUUID()
	//	h.Patiententrtime = time.Now()
	fmt.Println("h=",h)

//	hv = models.UpdatePatient(h)
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

func (this *StatusController) GetWarnStatus() { 
	logs.Debug("Post Line")
}
