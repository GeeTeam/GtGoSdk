package routers

import (
	"github.com/GeeTeam/GtGoSdk/demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register",&controllers.RegisterController{})
	beego.Router("/validate",&controllers.ValidateController{})
	beego.Router("/ajax_validate",&controllers.AjaxValidateController{})
}
