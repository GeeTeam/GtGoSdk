package controllers

import (
	"github.com/astaxie/beego"
	"github.com/GeeTeam/GtGoSdk"
)

var PrivateKey = beego.AppConfig.String("PrivateKey")
var CaptchaID = beego.AppConfig.String("CaptchaID")

type MainController struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller
}

type ValidateController struct {
	beego.Controller
}

type AjaxValidateController struct {
	beego.Controller
}

func (ctl *MainController) Get() {
	ctl.TplName = "index.html"
}

func (ctl *RegisterController)Get() {
	var userID = "test"
	gt := GtGoSdk.GeetestLib(PrivateKey, CaptchaID)
	status := gt.PreProcess(userID)
	ctl.SetSession(GtGoSdk.GT_STATUS_SESSION_KEY, status)
	ctl.SetSession("user_id", userID)
	responseStr := gt.GetResponseStr()
	ctl.Ctx.WriteString(responseStr)
}

func (ctl *ValidateController)Post() {
	var result bool
	var respstr string
	gt := GtGoSdk.GeetestLib(PrivateKey, CaptchaID)
	challenge := ctl.GetString(GtGoSdk.FN_CHALLENGE)
	validate := ctl.GetString(GtGoSdk.FN_VALIDATE)
	seccode := ctl.GetString(GtGoSdk.FN_SECCODE)
	status := ctl.GetSession(GtGoSdk.GT_STATUS_SESSION_KEY).(int)
	userID := ctl.GetSession("user_id").(string)
	if status == 0 {
		result = gt.FailbackValidate(challenge, validate, seccode)
	} else {
		result = gt.SuccessValidate(challenge, validate, seccode, userID)
	}
	if result {
		respstr = "<html><body><h1>登录成功</h1></body></html>"
	} else {
		respstr = "<html><body><h1>登录失败</h1></body></html>"
	}
	ctl.Ctx.WriteString(respstr)
}

func (ctl *AjaxValidateController)Post(){
	var result bool
	var jsondata = make(map[string]string)
	gt := GtGoSdk.GeetestLib(PrivateKey, CaptchaID)
	challenge := ctl.GetString(GtGoSdk.FN_CHALLENGE)
	validate := ctl.GetString(GtGoSdk.FN_VALIDATE)
	seccode := ctl.GetString(GtGoSdk.FN_SECCODE)
	status := ctl.GetSession(GtGoSdk.GT_STATUS_SESSION_KEY).(int)
	userID := ctl.GetSession("user_id").(string)
	if status == 0 {
		result = gt.FailbackValidate(challenge, validate, seccode)
	} else {
		result = gt.SuccessValidate(challenge, validate, seccode, userID)
	}
	if result {
		jsondata["status"]="success"
	} else {
		jsondata["status"]="fail"
	}
	ctl.Data["json"]= jsondata
	ctl.ServeJSON()
}