package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type PartController struct {
	xontrols.UICtrl
}

//parts/view/A10
func (c *PartController) Get() {
	c.Setup("partList", "Parts Inventory", true)

	result := []interface{}{}
	pagesize := c.FindParam("pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Stock.API", "part", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *PartController) GetCreate() {
	c.Setup("partCreate", "Parts Create", true)
}

func (c *PartController) GetView() {
	c.Setup("partView", "Parts View", true)

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		log.Println(err)
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Stock.API", "part", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
