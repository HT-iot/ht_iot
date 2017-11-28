package controllers

import (
	"github.com/astaxie/beego"
)

type Hospital struct {
	Hospital_id         string
	Hospital_zone       string
	Hospital_bed        int
	Patient_name        string
	Patient_sex         string
	Patient_id          string
	Patient_entr_time   string
	Patient_exit_time   string
	Patient_in_hospital bool
	Patient_urgent      bool
	Device_id           string
	Device_hospital_id  int
	Channel_id          string
	Meta                map[string]string
}

var Hospitalslice = make([]Hospital, 0, 100)

var Hospitalid int

type PconfigController struct {
	beego.Controller
}

func (this *PconfigController) Get() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplName = "pconfig.html"
	var h Hospital
	var err error

	op := this.Input().Get("op")
	id, err := this.GetInt("id")

	switch op {
	case "add":
		{
			h.Hospital_id = this.GetString("hospital_name")
			h.Hospital_zone = this.GetString("hospital_zone")
			h.Hospital_bed, err = this.GetInt("hospital_bed")
			if err != nil {
				return
			}
			h.Patient_name = this.GetString("p_name")
			h.Patient_sex = this.GetString("p_sex")
			h.Device_hospital_id, err = this.GetInt("p_device")
			h.Patient_urgent = false

			Hospitalslice = append(Hospitalslice, h)
			break
		}
	case "del":
		{
			Hospitalslice = append(Hospitalslice[:id], Hospitalslice[id+1:]...)
			//		Hospitalslice = append(Hospitalslice[:id], Hospitalslice[(id+1)]...)
			break
		}
	}
	//	this.Ctx.WriteString(fmt.Sprint(Hospitalslice))
	//	fmt.Println(Hospitalslice)
	this.Data["IsPconfig"] = true
	this.Data["Hospitalsilce"] = &Hospitalslice
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "pconfig.html"
}

func (this *PconfigController) Post() {

}
