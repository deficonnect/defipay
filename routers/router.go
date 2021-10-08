package routers

import (
	"tronbooth/controllers"

	"github.com/astaxie/beego" 
)

func init() {

	beego.Router("/form/:address", &controllers.MainController{}, "*:Form")
	beego.Router("/form", &controllers.MainController{}, "*:Form")

	beego.Router("/:address", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
}