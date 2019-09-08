package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"

	"github.com/louisevanderlith/service/controllers"
)

func Setup(e resins.Epoxi) {
	clients := &controllers.Clients{}
	parts := &controllers.Parts{}
	services := &controllers.Services{}
	e.JoinBundle("/stock", roletype.User, mix.Page, clients, parts, services)
}
