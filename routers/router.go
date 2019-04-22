package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
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
	beego.Router("/service", controllers.NewServiceCtrl(ctrlmap, theme))
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.User

	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/service", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
