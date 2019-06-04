package controllers

import (
	"log"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type ClientController struct {
	control.UIController
}

func NewClientCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *ClientController {
	result := &ClientController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ClientController) Get() {
	c.Setup("client", "Clients", true)
}

func (c *ClientController) GetCreate() {
	c.Setup("clientCreate", "Client Create", true)
}

func (c *ClientController) GetEdit() {
	c.Setup("clientEdit", "Client Edit", true)
}

func (c *ClientController) GetView() {
	c.Setup("clientView", "Client View", true)

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Entity.API", "info", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
