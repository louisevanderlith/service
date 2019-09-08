package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Services struct {
}

func (c *Services) Get(ctx context.Requester) (int, interface{}) {
	//c.Setup("serviceList", "Services", true)

	result := []interface{}{}
	pagesize := "A10"

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "service", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Services) Search(ctx context.Requester) (int, interface{}) {
	//c.Setup("serviceList", "Services", true)

	result := []interface{}{}
	pagesize := ctx.FindParam("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "service", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Services) Create(ctx context.Requester) (int, interface{}) {
	//c.Setup("serviceCreate", "Service Create", true)

	return http.StatusOK, nil
}

func (c *Services) View(ctx context.Requester) (int, interface{}) {
	//c.Setup("serviceView", "Service View", true)

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "service", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
