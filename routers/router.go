package routers

import (
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"

	"github.com/louisevanderlith/service/controllers"
)

func Setup(e resins.Epoxi) {
	//Default
	deftGroup := routing.NewInterfaceBundle("Home", roletype.User, &controllers.Home{})
	e.AddGroup(deftGroup)

	//Clients
	clientGroup := routing.NewInterfaceBundle("Clients", roletype.User, &controllers.Clients{})
	e.AddGroup(clientGroup)

	//Parts
	partGroup := routing.NewInterfaceBundle("Parts", roletype.User, &controllers.Parts{})
	e.AddGroup(partGroup)

	//Services
	servGroup := routing.NewInterfaceBundle("Services", roletype.Owner, &controllers.Services{})
	e.AddGroup(servGroup)
}
