package controllers

import (
	"log"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type PartController struct {
	control.UIController
}

func NewPartCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *PartController {
	result := &PartController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

//parts/view/A10
func (c *PartController) Get() {
	c.Setup("partList", "Parts Inventory", true)

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Stock.API", "part", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(result, nil)
}

func (c *PartController) GetCreate() {
	c.Setup("partCreate", "Parts Create", true)
}

func (c *PartController) GetView() {
	c.Setup("partView", "Parts View", true)

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Stock.API", "part", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
