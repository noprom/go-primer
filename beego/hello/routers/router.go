package routers

import (
	"github.com/noprom/go-primer/beego/hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
