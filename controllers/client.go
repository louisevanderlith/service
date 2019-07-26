package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type ClientController struct {
	xontrols.UICtrl
}

func (c *ClientController) Get() {
	c.Setup("clientList", "Clients", true)
}

func (c *ClientController) GetCreate() {
	c.Setup("clientCreate", "Client Create", true)
}

func (c *ClientController) GetEdit() {
	c.Setup("clientEdit", "Client Edit", true)
}

func (c *ClientController) GetView() {
	c.Setup("clientView", "Client View", true)

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Entity.API", "info", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
