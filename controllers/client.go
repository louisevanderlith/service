package controllers

import (
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
	c.Setup("clientCreate", "Client", true)
}

func (c *ClientController) GetEdit() {
	c.Setup("clientEdit", "Client", true)
}

func (c *ClientController) GetView() {
	c.Setup("clientView", "Client", true)
}
