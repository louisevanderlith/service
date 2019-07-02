package controllers

import (
	"log"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type ServiceController struct {
	control.UIController
}

func NewServiceCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *ServiceController {
	result := &ServiceController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ServiceController) Get() {
	c.Setup("serviceList", "Services", true)

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Stock.API", "service", "all", pagesize)

	c.Serve(result, err)
}

func (c *ServiceController) GetCreate() {
	c.Setup("serviceCreate", "Service Create", true)
}

func (c *ServiceController) GetView() {
	c.Setup("serviceView", "Service View", true)

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Stock.API", "service", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
