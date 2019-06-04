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
	beego.Router("/service", srvCtrl)

	clientCtrl := controllers.NewClientCtrl(ctrlmap, theme)
	beego.Router("/client/edit/:key", clientCtrl, "get:GetEdit")
	beego.Router("/client/create", clientCtrl, "get:GetCreate")
	beego.Router("/client/:key", clientCtrl, "get:GetView")

	//beego.Router("/service", srvCtrl)
	//beego.Router("/service", srvCtrl)
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.User

	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/service", emptyMap)
	ctrlmap.Add("/client", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
