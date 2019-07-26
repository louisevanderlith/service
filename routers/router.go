package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"

	"github.com/louisevanderlith/service/controllers"
)

func Setup(e *droxolite.Epoxy) {
	//Default
	deftCtrl := &controllers.DefaultController{}
	deftGroup := droxolite.NewRouteGroup("", deftCtrl)
	deftGroup.AddRoute("/", "GET", roletype.User, deftCtrl.Get)
	e.AddGroup(deftGroup)

	//Clients
	clientCtrl := &controllers.ClientController{}
	clientGroup := droxolite.NewRouteGroup("clients", clientCtrl)
	clientGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, clientCtrl.GetView)
	clientGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, clientCtrl.GetEdit)
	clientGroup.AddRoute("/create", "GET", roletype.User, clientCtrl.GetCreate)
	e.AddGroup(clientGroup)

	//Parts
	partCtrl := &controllers.PartController{}
	partGroup := droxolite.NewRouteGroup("parts", partCtrl)
	partGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, partCtrl.GetView)
	partGroup.AddRoute("/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, partCtrl.Get)
	partGroup.AddRoute("/create", "GET", roletype.User, partCtrl.GetCreate)
	e.AddGroup(partGroup)

	//Services
	servCtrl := &controllers.ServiceController{}
	servGroup := droxolite.NewRouteGroup("services", servCtrl)
	servGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, servCtrl.GetView)
	servGroup.AddRoute("/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, servCtrl.Get)
	servGroup.AddRoute("/create", "GET", roletype.User, servCtrl.GetCreate)
	e.AddGroup(servGroup)
	/* := EnableFilter(s)

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
	beego.Router("/client/:key", clientCtrl, "get:GetView")*/
}

/*
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
*/
