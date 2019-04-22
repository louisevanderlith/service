package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type DefaultController struct {
	control.UIController
}

func NewDefaultCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *DefaultController {
	result := &DefaultController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *DefaultController) Get() {
	c.Setup("default", "Service Home", false)
}
