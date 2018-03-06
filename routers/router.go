package routers

import (
	"github.com/lvrenhui/mock_server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// admin.Run()
	beego.Router("/", &controllers.MockController{}, "*:List")
	beego.Router("/config/list", &controllers.MockController{}, "*:List")
	beego.Router("/config/add", &controllers.MockController{}, "get:AddPage")
	beego.Router("/config/add", &controllers.MockController{}, "post:Add")
	beego.Router("/config/mod", &controllers.MockController{}, "get:UpdatePage")
	beego.Router("/config/mod", &controllers.MockController{}, "post:Update")
	beego.Router("/config/mod_check", &controllers.MockController{}, "post:UpdateCheck")
	beego.Router("/config/del", &controllers.MockController{}, "post:Del")
	beego.Router("/config/get", &controllers.MockController{}, "post:Get")
}
