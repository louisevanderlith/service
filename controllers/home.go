package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Home struct {
}

func (c *Home) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("default", "Service Home", true)

	return http.StatusOK, nil
}
