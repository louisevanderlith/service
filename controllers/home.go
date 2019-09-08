package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

func Index(ctx context.Requester) (int, interface{}) {
	//c.Setup("default", "Service Home", true)

	return http.StatusOK, nil
}
