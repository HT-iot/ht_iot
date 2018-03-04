package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"ht_iot/models"
	"io"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"golang.org/x/crypto/bcrypt"
)

const cost int = 10

type bcryptHasher struct{}
func (c *LoginController) Hash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (c *LoginController) Compare(plain, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
	isExit := c.GetString("exit")
	fmt.Println("Login isExit =", isExit)
	if isExit == "true" {
		c.Ctx.SetCookie("name", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Ctx.SetCookie("hname", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
}

func (c *LoginController) Post() {
	log := logs.GetBeeLogger();

	c.TplName = "login.html"
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	hname := c.Input().Get("hospname")

	autoLogin := c.Input().Get("autoLogin") == "on"

	//p := models.User{Email: uname, Password: pwd}
	p := models.User{Email: uname}
	q, err := models.GetAllUers(p)

	fmt.Println("q, err =", q, err)

	if len(q) > 0 {
		log.Debug("user input passwd is :" + pwd)
		log.Debug("db passwd is :" + q[0].Password)

		err = c.Compare(pwd, q[0].Password);
		if  err != nil {
			log.Debug("passwd compare failure !")
		}else{
			log.Debug("passwd compare success !")
		}
	}	

	if err == nil {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("name", q[0].Email, maxAge, "/")
		h := sha1.New()
		io.WriteString(h, q[0].Email+q[0].Password)
		pwd = base64.StdEncoding.EncodeToString(h.Sum(nil))
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/")

		hname0 := base64.StdEncoding.EncodeToString([]byte(hname))
		c.Ctx.SetCookie("hname", hname0, maxAge, "/")

		/* TODO add session support
		c.SetSession(uname, pwd)
		*/
		IsLogin = true
		c.Data["ISLogin"] = IsLogin

		c.TplName = "home.html"
	} else {
		IsLogin = false
		c.Data["ISLogin"] = IsLogin
		c.Redirect("/login", 301)
	}
	return
}

func checkAccount(ctx *context.Context) bool {
	log := logs.GetBeeLogger();

	ck, err := ctx.Request.Cookie("name")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	ck, err = ctx.Request.Cookie("hname")
	if err != nil {
		return false
	}
	hname, err := base64.StdEncoding.DecodeString(ck.Value)
	if err != nil {
		fmt.Println(err)
	}
	Hospital = string(hname)
	fmt.Println(Hospital)

	p := models.User{Email: uname}
	q, err := models.GetAllUers(p)

	if len(q) < 1 {
		logs.Debug("Query user table failure")
		IsLogin = false
		return false
	}

	log.Debug("user input passwd is :" + pwd)
	log.Debug("db passwd is :" + q[0].Password)

	h := sha1.New()
	io.WriteString(h, q[0].Email+q[0].Password)
	if pwd == base64.StdEncoding.EncodeToString(h.Sum(nil)) {
		logs.Debug("compare pwd success")
		IsLogin = true
		return true
	}

	logs.Debug("Cookie password check failure")
	IsLogin = false
	return false

}
