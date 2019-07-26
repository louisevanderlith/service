package controllers

import "github.com/louisevanderlith/droxolite/xontrols"

type DefaultController struct {
	xontrols.UICtrl
}

func (c *DefaultController) Get() {
	c.Setup("default", "Service Home", true)
}
