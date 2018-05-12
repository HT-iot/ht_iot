package controllers

import (
	"fmt"

	"ht_iot/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
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
	s.TplName = "patient_status.html"
	flag := checkAccount(s.Ctx)
	s.Data["ISLogin"] = flag
//	s.Data["Hospital"] = Hospital
	if !flag {
		s.Redirect("/login", 302)
		return
	}
	_,_ = Calculatepara()
}

/*GetStatus reponse the API query */
func (s *StatusController) GetStatus() {
	logs.Debug("GetStatus")
//	var hPatients []models.HospitalPatientInfo
	var data models.DataTable
	_ = checkAccount(s.Ctx)
	data.Data,_ = Calculatepara()
//	logs.Debug("Status Patients:", data.Data)
	s.Data["json"] = &data
	s.ServeJSON()
}

/*PostStatus reponse the API query */
func (s *StatusController) PostStatus() {
	logs.Debug("Input the data in Patient Status")
	col := [][2]string{
		{"reporttime","上报时间"},          
		{"hospitalname","医院名称"},
		{"hospitalzone", "病区号"}, 
		{"hospitalbed", "病床号"}, 
		{"patientname", "姓名"},
		{"hospitaldeviceid", "终端号"}, 
		{"puls", "脉搏"}, 
		{"oxgen", "血氧"},
		{"pressurelow", "舒张压"},
		{"pressurehigh", "收缩压"},
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
	}

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
	_,Para := Calculatepara()
	this.Data["json"] = &Para
	this.ServeJSON()
}



func (this *StatusController) GetInfo() {
	logs.Debug("Hospital warn config information to Config table")
//	var Mystruct models.HospitalInfoTable
 	hname := Hospital
	hzone := HospZone

	pHC,_ = models.GetWarnInfo(hname, hzone)	

	logs.Debug("HospitalInfoConfig:", pHC)
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

	logs.Debug("h=",h)
	var Getstruct In

	if (!(models.UpdateWarnInfo(h))){
//		fmt.Println("updated failed")
		if err:= models.InsertWarnInfo(h); err !=nil{
			Getstruct.Info = "添加失败"
			Getstruct.Succ = "fail"
		}
	}
	Getstruct.Info = "更新成功"
	Getstruct.Succ = "Succ"

	this.Data["json"] = &Getstruct
	this.ServeJSON()
}


func Calculatepara()(Patients []models.PatientInfo, Para Warndata){

	logs.Debug("Get Latest status/Information of Patient  to display")
	var hPatients []models.HospitalPatientInfo
	
	Para.Totalpatient = 0 
	Para.Totaloxgen =0
	Para.Totalpuls = 0
	Para.Totalurgent =0
	Para.Totaluses = 0 
	Para.Totalunuses = 0 
	Para.Totalpress = 0 

	name := Hospital
	zone := HospZone
	logs.Debug("Calculatepara hospital teste =================", name, zone)
//	hzone := string("")

	pHC,_ = models.GetWarnInfo(name, zone)	
	
	// Get all in-hospital patients.
	if len(name) >0  {
		hp := models.HospitalPatientInfo{Hospitalname: name, Hospitalzone:zone}
		hPatients, _ = models.GetPatient(hp)
	} else {
		hPatients, _ = models.GetAllPatient()
	}

	if len(hPatients) == 0 {
		logs.Error("Can't get patient info from DB")
	}

//Calcualte the number of real uses;
	var i = 0
	for _, p := range hPatients {
		if (len(p.Patientname)>0) {i++;}
	}
	Para.Totalpatient  = i

	Patients = make([]models.PatientInfo, i)

//Move the real use patient information to Patients
	i = 0
	for _, p := range hPatients {
		if (len(p.Patientname)>0){
			Patients[i].Hospitalname = p.Hospitalname
			Patients[i].Hospitalzone = p.Hospitalzone
			Patients[i].Hospitalbed = p.Hospitalbed
			Patients[i].Patientname = p.Patientname
			Patients[i].Hospitaldeviceid = p.Hospitaldeviceid
			Patients[i].Channelid = p.Channelid
			Patients[i].Deviceid = p.Deviceid
		
			err := models.GetMsg(&Patients[i])   //assemble the data
			if err != nil {
				logs.Error("Get the status failure!")
			}
//			logs.Debug(Patients[i])
		i++;
		}
	}

//计算指标
//	
	i=0
	for i, _ := range Patients {
		switch{
		case ((Patients[i].Runstatus)=="在线"): 
			{
				Para.Totaluses++;
//正常带				
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
			}
		case ((Patients[i].Runstatus)=="在线危急"):
			{
				Para.Totalurgent++;
			}

		case ((Patients[i].Runstatus)=="离线"):
			{
				Para.Totalunuses++;
			}
		case ((Patients[i].Runstatus)=="数据不完整"):
			{

			} 
		}
	}

//	logs.Debug("MSG Patients",Patients)
	return Patients, Para
}

func (this *StatusController) GetOneData() { 
	logs.Debug("Get All information on One days about the patient")
//	var data models.DataTable
	
	t1:=time.Now()
	Deviceid := this.GetString("deviceid")
	Hospitalname := this.GetString("hospitalname")
	Patientname := this.GetString("patientname")
	days := this.GetString("days")
	
	var day int
	if (days == "一"){
		day=1
	}else{
		if(days == "五"){
			day=5
		}else{
			day = 3650
		}
	}
	
	fmt.Print("t1=",Hospitalname, Patientname, t1.Unix())

	t2 := t1.AddDate(0,0,(0-day))
	t2time := float64(t2.Unix())
	logs.Debug("t2time %s", t2time)
//	var patients []models.PatientInfos
	var data models.DataTable
//	data.Data,_ = models.GetDayInfo(Hospitalname,Patientname,t2time)
	data.Data = models.GetDayInfo(Deviceid,Hospitalname,Patientname,t2time)
//	data.Data[0].Patientname ="11111"
	this.Data["json"] = &data
	this.ServeJSON()
}

