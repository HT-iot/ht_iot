package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

var Hospitalslice []Out

type Out struct {
	Hospitalname     string `json:"hospitalname"`
	Hospitalzone     string `json:"hospitalzone"`
	Hospitalbed      string `json:"hospitalbed"`
	Patientname      string `json:"patientname"`
	Hospitaldeviceid string `json:"hospitaldeviceid"`
}

type Jsonout struct {
	Draw            int   `json:"draw"`
	Recordstotal    int   `json:"recordstotal"`
	Recordsfiltered int   `json:"recordsfiltered"`
	Data            []Out `json:"data"`
}

/* PconfigController...*/
type PconfigController struct {
	beego.Controller
}

/*
func (this *PconfigController) Prepare() {
	this.TplName = "pconfig.html"
	this.Data["IsPconfig"] = true

	flag := checkAccount(this.Ctx)
	this.Data["ISLogin"] = flag
	if !flag {
		this.Redirect("/login", 302)
		return
	}
	this.Data["Hospitalsilce"] = &Hospitalslice
	this.TplName = "pconfig.html"
}
*/
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
	var h Out
	{
		h.Hospitalname = this.GetString("hospital_name")
		h.Hospitalzone = this.GetString("hospital_zone")
		h.Hospitalbed = this.GetString("hospital_bed")
		h.Patientname = this.GetString("p_name")
		//		h.Patientsex = this.GetString("p_sex")
		h.Hospitaldeviceid = this.GetString("p_device")
	}
	Hospitalslice = append(Hospitalslice, h)

	//	this.Data["Hospitalsilce"] = &Hospitalslice

}

func (this *PconfigController) Post() {

	fmt.Println("Post")

	var Mystruct Jsonout

	Mystruct.Draw = len(Hospitalslice)
	Mystruct.Recordstotal = len(Hospitalslice)
	Mystruct.Recordsfiltered = -1
	Mystruct.Data = append(Mystruct.Data, Hospitalslice...)

	/*
		Mystruct.Data = append(Mystruct.Data, Out{
			Name:       "Ashton Cox",
			Position:   "Junior Technical Author",
			Salary:     "$86,0",
			StartDate:  "2009/01/12",
			Salary2:    "$86000",
			StartDate2: "2010/01/12",
			StartDate3: "2011/01/12",
			Office:     "San Francisco",
			Extn:       "1562",
		})
		Mystruct.Data = append(Mystruct.Data, Out{
			Name:       "Ashton 2",
			Position:   "Junior Technical Author",
			Salary:     "$86000",
			StartDate:  "2009/01/12",
			Salary2:    "$86,000",
			StartDate2: "2010/01/12",
			StartDate3: "2011/01/12",
			Office:     "San Francisco",
			Extn:       "1562",
		})

		Mystruct.Data = append(Mystruct.Data, Out{
			Name:       "Ashton 3",
			Position:   "Junior Technical Author",
			Salary:     "$8600",
			StartDate:  "2009/01/12",
			Salary2:    "$86,000",
			StartDate2: "2010/01/12",
			StartDate3: "2011/01/12",
			Office:     "San Francisco",
			Extn:       "1562",
		})


			: {
				Name:       "Ashton Cox",
				Position:   "Junior Technical Author",
				Salary:     "$86,000",
				StartDate:  "2009/01/12",
				Salary2:    "$86,000",
				StartDate2: "2010/01/12",
				StartDate3: "2011/01/12",
				Office:     "San Francisco",
				Extn:       "1562",
			},
			{
				Name:       "Ashton Cox",
				Position:   "Junior Technical Author",
				Salary:     "$86,000",
				StartDate:  "2009/01/12",
				Salary2:    "$86,000",
				StartDate2: "2009/01/12",
				StartDate3: "2009/01/12",
				Office:     "San Francisco",
				Extn:       "15",
			}}
	*/ /*


			var h models.HospitalPatientInfo

		lineNO, _ := this.GetInt("id")
		fmt.Println("Line:", lineNO)

		Hospitalslice[lineNO].Id = gocql.TimeUUID()
		h = Hospitalslice[lineNO]
		//		h.Patiententrtime = time.Now()
		hv := falseh.Channelid, h.Deviceid, hv = models.GetPaientIDs(Hospitalslice[lineNO].Hospitalname, Hospitalslice[lineNO].Hospitaldeviceid)
		if hv {
			_ = models.InsertPatient(h)
			Hospitalslice = append(Hospitalslice[:lineNO], Hospitalslice[lineNO+1:]...)

			mystruct.Succ = "添加成功"
			mystruct.Refresh = "44"
			this.Data["Hospitalsilce"] = &Hospitalslice
		} else {
			mystruct.Succ = "添加失败, 无该医院终端ID"
			mystruct.Refresh = "00"
		}
		//	mystruct = Out{Succ: str, Refresh: Out.Refresh}
	*/
	this.Data["json"] = &Mystruct

	this.ServeJSON()

	//	this.Data["Hospitalsilce"] = &Hospitalslice
	//	this.TplName = "pconfig.html"
}
