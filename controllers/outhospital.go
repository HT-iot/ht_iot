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


type Column struct {
	checkfield [2]string
	headfield [][2]string 
}

func (this *OuthospitalController) GetPat() {
	logs.Debug("Input the data in Pconfig")
/*	
	var col [1]Column
	fmt.Println(col);

	col[0].checkfield = [2]string {"hospitalname", "医院名称"}
	col[0].headfield = [][2]string {{"hospitalname","医院名称"}}


	col[1].checkfield[0] = "hospitalzone"
	col[1].checkfield[1] = "病区号"
	col[1].headfield[0][0] = "hospitalzone"
	col[1].headfield[0][1] = "病区号"

	col[2].checkfield[0] = "hospitalbed"
	col[2].checkfield[1] = "病床号"
	col[2].headfield[0][0] = "hospitalbed"
	col[2].headfield[0][1] = "病床号"

	col[3].checkfield[0] = "patientname"
	col[3].checkfield[1] = "姓名"
	col[3].headfield[0][0] = "patientname"
	col[3].headfield[0][1] = "姓名"

	col[4].checkfield[0] = "hospitaldeviceid"
	col[4].checkfield[1] = "终端号"
	col[4].headfield[0][0] = "hospitaldeviceid"
	col[4].headfield[0][1] = "终端号"

	col[5].checkfield[0] = "puls"
	col[5].checkfield[1] = "脉搏"
	col[5].headfield[0][0] = "puls"
	col[5].headfield[0][1] = "脉搏"

	col[6].checkfield[0] = "oxgen"
	col[6].checkfield[1] = "血氧"
	col[6].headfield[0][0] = "oxgen"
	col[6].headfield[0][1] = "血氧"

	col[7].checkfield[0] = "pressure"
	col[7].checkfield[1] = "血压"
	col[7].headfield[0][0] = "pressurehigh"
	col[7].headfield[0][1] = "收缩压"
	col[7].headfield[1][0] = "pressurelow"
	col[7].headfield[1][1] = "舒张压"

	col[8].checkfield[0] = "poistion"
	col[8].checkfield[1] = "位置"
	col[8].headfield[0][0] = "longitude"
	col[8].headfield[0][1] = "经度"
	col[8].headfield[1][0] = "latitude"
	col[8].headfield[1][1] = "纬度"


	col := [][2]string{
		{"hospitalname","医院名称"},
		{"hospitalzone", "病区号"}, 
		{"hospitalbed", "病床号"}, 
		{"patientname", "姓名"},
		{"hospitaldeviceid", "终端号"}, 
		{"puls", "脉搏"}, 
		{"oxgen", "血氧"},
		{"pressure", "血压"},
		{"position", "位置"},
	
	}
*/
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


	//col := []string{"hospitalzone", "hospitalbed", "patientname", "hospitaldeviceid", "puls", "pressure", "oxgen"}
	//	col := []string{"病区号", "病床号", "姓名", "终端号", "脉搏", "血压", "血氧"}

	this.Data["json"] = &col
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
