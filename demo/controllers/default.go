package controllers

import (
	"github.com/astaxie/beego"
	"github.com/GeeTeam/GtGoSdk"
	"log"
)

var GtPrivateKey = beego.AppConfig.String("GtPrivateKey")
var GtCaptchaID = beego.AppConfig.String("GtCaptchaID")

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
	gt := GtGoSdk.GeetestLib(GtPrivateKey, GtCaptchaID)
	gt.PreProcess(userID)
	responseMap := gt.GetResponseMap()
	ctl.Data["json"]=responseMap
	ctl.ServeJSON()
}

func (ctl *ValidateController)Post() {
	var result bool
	var respstr string
	gt := GtGoSdk.GeetestLib(GtPrivateKey, GtCaptchaID)
	challenge := ctl.GetString(GtGoSdk.FN_CHALLENGE)
	validate := ctl.GetString(GtGoSdk.FN_VALIDATE)
	seccode := ctl.GetString(GtGoSdk.FN_SECCODE)
	status,err := ctl.GetInt(GtGoSdk.GT_STATUS_SESSION_KEY)
	if err != nil{
		log.Println(err)
	}
	userID := ctl.GetString("user_id")
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
	gt := GtGoSdk.GeetestLib(GtPrivateKey, GtCaptchaID)
	challenge := ctl.GetString(GtGoSdk.FN_CHALLENGE)
	validate := ctl.GetString(GtGoSdk.FN_VALIDATE)
	seccode := ctl.GetString(GtGoSdk.FN_SECCODE)
	status,err := ctl.GetInt(GtGoSdk.GT_STATUS_SESSION_KEY)
	userID := ctl.GetString("user_id")
	if err != nil{
		log.Println(err)
	}
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