package main

import (
	"reflect"
	"strings"

	"github.com/lvrenhui/mock_server/controllers"
	_ "github.com/lvrenhui/mock_server/routers"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func main() {
	RouteFilter := func(ctx *context.Context) {
		uri := ctx.Request.RequestURI
		if len(uri) > 1 && uri != "/favicon.ico" && !strings.HasPrefix(uri, "/config") {
			ctx.Input.RunController = reflect.TypeOf(controllers.MockController{})
			ctx.Input.RunMethod = "Mock"
		}
		// if strings.HasPrefix(uri, "/config") || uri == "/" {
		// 	ctx.Input.RunController = reflect.TypeOf(controllers.MockController{})
		// 	if uri == "/" || strings.HasSuffix(uri, "list") {
		// 		ctx.Input.RunMethod = "List"
		// 	} else if strings.HasSuffix(uri, "add") {
		// 		if ctx.Input.Method() == "POST" {
		// 			ctx.Input.RunMethod = "Add"
		// 		} else {
		// 			ctx.Input.RunMethod = "AddPage"
		// 		}
		// 	}
		// } else {

		// }

	}
	// fmt.Println(RouteFilter)

	//通过过滤器拦截url
	beego.InsertFilter("/*", beego.BeforeRouter, RouteFilter)
	beego.Run()
}
