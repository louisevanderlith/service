package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"

	//"github.com/louisevanderlith/secure/core/roletype"
	"github.com/louisevanderlith/service/controllers"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	siteName := beego.AppConfig.String("defaultsite")
	theme, err := mango.GetDefaultTheme(ctrlmap.GetInstanceID(), siteName)

	if err != nil {
		panic(err)
	}

	beego.Router("/", controllers.NewDefaultCtrl(ctrlmap, theme))

	srvCtrl := controllers.NewServiceCtrl(ctrlmap, theme)
	beego.Router("/services/view/:pagesize", srvCtrl, "get:Get")
	beego.Router("/services/create", srvCtrl, "get:GetCreate")
	beego.Router("/service/:key", srvCtrl, "get:GetView")

	partCtrl := controllers.NewPartCtrl(ctrlmap, theme)
	beego.Router("/parts/view/:pagesize", partCtrl, "get:Get")
	beego.Router("/parts/create", partCtrl, "get:GetCreate")
	beego.Router("/part/:key", partCtrl, "get:GetView")

	clientCtrl := controllers.NewClientCtrl(ctrlmap, theme)
	beego.Router("/client/edit/:key", clientCtrl, "get:GetEdit")
	beego.Router("/client/create", clientCtrl, "get:GetCreate")
	beego.Router("/client/:key", clientCtrl, "get:GetView")
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.User

	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/service", emptyMap)
	ctrlmap.Add("/part", emptyMap)
	ctrlmap.Add("/client", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
