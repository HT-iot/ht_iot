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

type Warndata struct {
	Totalpatient	int `json:"totalpatient"`
	Totaluses		int `json:"totaluses"`
	Totalunuses		int `json:"totalunuses"`
	Totalurgent		int `json:"totalurgent"`
	Totalpuls		int `json:"totalpuls"`
	Totaloxgen		int `json:"totaloxgen"`
	Totalpress		int `json:"totalpress"`
}

var pHC models.HospitalInfoConfig
//var Para Warndata
//var hPatients []models.HospitalPatientInfo

/*Get reponse the status query */
func (s *StatusController) Get() {
	s.Data["IsSconfig"] = true

	flag := checkAccount(s.Ctx)
	s.Data["ISLogin"] = flag
	if !flag {
		s.Redirect("/login", 302)
		return
	}
	
	_,_ = Calculatepara()
	s.TplName = "patient_status.html"
}

/*GetStatus reponse the API query */
func (s *StatusController) GetStatus() {
	logs.Debug("GetStatus")
//	var hPatients []models.HospitalPatientInfo
	var data models.DataTable

//	name := Hospital
/*	name := s.GetString("hospital_name")
//	zone := s.GetString("hospital_zone")
//	pName := s.GetString("p_name")

	if name != "" || zone != "" || pName != "" {
		hp := models.HospitalPatientInfo{
			//Hospitalname: name,
			//Hospitalzone: zone,
			Patientname: pName}
		hPatients, _ = models.GetPatient(hp)
	} else {
		hPatients, _ = models.GetAllPatient()
	}
*/
/*
	if name != "" {
		hp := models.HospitalPatientInfo{
			Hospitalname: name}
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
//计算指标
	Para.Totalpatient = 0 
	Para.Totaloxgen =0
	Para.Totalpuls = 0
	Para.Totalurgent =0
	Para.Totaluses = 0 
	Para.Totalunuses = 0 
	Para.Totalpress = 0 

	
	Para.Totalpatient  = len(Patients)
	
	for i, _ := range Patients {
		if((Patients[i].Patientname != "") && (Patients[i].Hospitalbed != "")&& (Patients[i].Hospitaldeviceid != "")){
			if ((Patients[i].Puls>0)||(Patients[i].Oxgen>0)||(Patients[i].Pressurehigh>0)||(Patients[i].Pressurelow>0)){
					Para.Totaluses++;
				if ((Patients[i].Puls <= pHC.Pulsmin) ||(Patients[i].Puls > pHC.Pulsmax)) {
					Para.Totalpuls++;
				}
				if ((Patients[i].Oxgen <= pHC.Oxgenmin) ||(Patients[i].Oxgen > pHC.Oxgenmax)) {
					Para.Totaloxgen++;
				}
				if ((Patients[i].Pressurehigh >= pHC.Pressurehighmax) ||(Patients[i].Pressurehigh <= pHC.Pressurehighmin)||
				(Patients[i].Pressurelow >= pHC.Pressurelowmax)||(Patients[i].Pressurelow <= pHC.Pressurelowmin)){
					Para.Totalpress++;
				}
			} else{
				Para.Totalunuses++;
			}
//未判断有体感，无信号的Totalurgent.				
		} else{
//			Totalurgent++;
		}
	}
	
//	Totalunuses = Totalpatient-Totalurgent-Totaluses
fmt.Println("ParaAA=",Para)

*/

	data.Data,_ = Calculatepara()
	fmt.Println("Patients",data.Data)
	s.Data["json"] = &data
	s.ServeJSON()
}

/*PostStatus reponse the API query */
func (s *StatusController) PostStatus() {
	logs.Debug("Input the data in Patient Status")
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
		{"ops", "操作"},
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
		h.Pulsmin,_ = this.GetFloat("pulsmin")
		h.Pulsmax,_ = this.GetFloat("pulsmax")
		h.Oxgenmin,_ = this.GetFloat("oxgenmin")
		h.Oxgenmax,_ = this.GetFloat("oxgenmax")
		h.Pressurelowmin,_ = this.GetFloat("pressurelowmin")
		h.Pressurelowmax,_ = this.GetFloat("pressurelowmax")
		h.Pressurehighmin,_ = this.GetFloat("pressurehighmin")
		h.Pressurehighmax,_ = this.GetFloat("pressurehighmax")
		h.Monitoraddress = this.GetString("monitoraddress")
		h.Monitorradius,_ = this.GetFloat("monitorradius")
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
	logs.Debug("Send the warn table for the warning diagram")
/*	var out Warndata
/
	out.Totalpatient = Totalpatient
	out.Totaloxgen = Totaloxgen
	out.Totalpress = Totalpress
	out.Totalpuls = Totalpress
	out.Totalunuses = Totalunuses
	out.Totalurgent = Totalurgent
	out.Totaluses = Totaluses
*/
	_,Para := Calculatepara()
	fmt.Println("Para==",Para)
	this.Data["json"] = &Para
	this.ServeJSON()

}



func (this *StatusController) GetInfo() {
	logs.Debug("Hospital warn config information to Config table")
//	var Mystruct models.HospitalInfoTable
 	hname := Hospital
	hzone := string("")

	pHC,_ = models.GetWarnInfo(hname, hzone)	

	fmt.Println("HospitalInfoConfig:", pHC)
	this.Data["json"] = &pHC
	this.ServeJSON()
	
}

func (this *StatusController) PostInfo() {
	logs.Debug("Get warn information from Device table")

	var h  models.HospitalInfoConfig

	//	var ob models.HospitalInfoTable;
	h.Hospitalname = this.Input().Get("Hospitalname")
	h.Hospitalzone = this.Input().Get("Hospitalzone")
	h.Pulsmin,_ = this.GetFloat("Pulsmin")
	h.Pulsmax,_ = this.GetFloat("Pulsmax")

	h.Oxgenmin,_ = this.GetFloat("Oxgenmin")
	h.Oxgenmax,_ = this.GetFloat("Oxgenmax")
	h.Pressurelowmin,_ = this.GetFloat("Pressurelowmin")
	h.Pressurelowmax,_ = this.GetFloat("Pressurelowmax")

	h.Pressurehighmin,_ = this.GetFloat("Pressurehighmin")
	h.Pressurehighmax,_ = this.GetFloat("Pressurehighmax")
	h.Monitoraddress = this.Input().Get("Monitoraddress")
	h.Monitorradius,_ = this.GetFloat("Monitorradius")

	fmt.Println("h=",h)

	if (!(models.UpdateWarnInfo(h))){
		fmt.Println("updated failed")
		models.InsertWarnInfo(h)
	}


var Getstruct In
Getstruct.Info = "无该医院终端ID, 添加失败"
Getstruct.Succ = "fail"

	this.Data["json"] = &Getstruct
	this.ServeJSON()
}


func Calculatepara()(Patients []models.PatientInfo, Para Warndata){

	logs.Debug("GetStatus of information")
	var hPatients []models.HospitalPatientInfo
	name := Hospital
	hzone := string("")
	pHC,_ = models.GetWarnInfo(name, hzone)	
	
	if name != "" {
		hp := models.HospitalPatientInfo{Hospitalname: name}
		hPatients, _ = models.GetPatient(hp)
	} else {
		hPatients, _ = models.GetAllPatient()
	}

	if len(hPatients) == 0 {
		logs.Error("Can't get patient info from DB")
	}

	Patients = make([]models.PatientInfo, len(hPatients))
//	var data models.DataTable

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
//计算指标
	Para.Totalpatient = 0 
	Para.Totaloxgen =0
	Para.Totalpuls = 0
	Para.Totalurgent =0
	Para.Totaluses = 0 
	Para.Totalunuses = 0 
	Para.Totalpress = 0 
	
	Para.Totalpatient  = len(Patients)
	
	for i, _ := range Patients {
		if((Patients[i].Patientname != "") && (Patients[i].Hospitalbed != "")&& (Patients[i].Hospitaldeviceid != "")){
			if ((Patients[i].Puls>0)||(Patients[i].Oxgen>0)||(Patients[i].Pressurehigh>0)||(Patients[i].Pressurelow>0)){
					Para.Totaluses++;
				if ((Patients[i].Puls <= pHC.Pulsmin) ||(Patients[i].Puls > pHC.Pulsmax)) {
					Para.Totalpuls++;
				}
				if ((Patients[i].Oxgen <= pHC.Oxgenmin) ||(Patients[i].Oxgen > pHC.Oxgenmax)) {
					Para.Totaloxgen++;
				}
				if ((Patients[i].Pressurehigh >= pHC.Pressurehighmax) ||(Patients[i].Pressurehigh <= pHC.Pressurehighmin)||
				(Patients[i].Pressurelow >= pHC.Pressurelowmax)||(Patients[i].Pressurelow <= pHC.Pressurelowmin)){
					Para.Totalpress++;
				}
			} else{
				Para.Totalunuses++;
			}
//未判断有体感，无信号的Totalurgent.				
		} else{
//			Totalurgent++;
		}
	}
	return Patients, Para
}
