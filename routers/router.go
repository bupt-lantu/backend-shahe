// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"bupt_tour/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/place",
			beego.NSInclude(&controllers.PlaceController{}),
		),
		beego.NSNamespace("/placetype",
			beego.NSInclude(&controllers.PlaceTypeController{}),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(&controllers.AuthController{}),
		),
	)
	beego.AddNamespace(ns)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//AllowAllOrigins: ,
		AllowOrigins:     []string{"http://localhost:4200", "http://10.3.244.81:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Controlet g:ale_fix_on_save = 1l-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/v1/*", beego.BeforeRouter, FilterUser)
}

func FilterUser(ctx *context.Context) {
	if ctx.Request.Method != "GET" && ctx.Request.URL.String() != "/v1/login" {
		_, ok := ctx.Input.Session("id").(int)
		if !ok {
			ctx.ResponseWriter.WriteHeader(401)
			ctx.Abort(401, "未登录")
		}
	}
}
