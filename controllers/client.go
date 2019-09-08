package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Clients struct {
}

func (c *Clients) Get(ctx context.Requester) (int, interface{}) {
	//c.Setup("clientList", "Clients", true)

	return http.StatusOK, nil
}

func (c *Clients) Search(ctx context.Requester) (int, interface{}) {
	//c.Setup("clientList", "Clients", true)

	return http.StatusOK, nil
}

func (c *Clients) View(ctx context.Requester) (int, interface{}) {
	//c.Setup("clientEdit", "Client Edit", true)

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Entity.API", "info", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Clients) Create(ctx context.Requester) (int, interface{}) {
	//c.Setup("clientCreate", "Client Create", true)
	return http.StatusOK, nil
}
