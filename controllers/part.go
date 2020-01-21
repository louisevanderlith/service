package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Parts struct {
}

func (c *Parts) Get(c *gin.Context) {
	//c.Setup("partList", "Parts Inventory", true)

	result := []interface{}{}
	pagesize := "A10"

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "part", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

//parts/view/A10
func (c *Parts) Search(c *gin.Context) {
	//c.Setup("partList", "Parts Inventory", true)

	result := []interface{}{}
	pagesize := c.Param("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "part", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Parts) Create(c *gin.Context) {
	//c.Setup("partCreate", "Parts Create", true)

	return http.StatusOK, nil
}

func (c *Parts) View(c *gin.Context) {
	//c.Setup("partView", "Parts View", true)

	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "part", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
