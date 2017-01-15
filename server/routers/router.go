package routers

import (
	"github.com/astaxie/beego"
	"meiya/server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/uploadPic", &controllers.UploadController{})
}
