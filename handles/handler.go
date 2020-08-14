package handles

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(clnt, scrt, securityUrl, authorityUrl string) http.Handler {
	/*clients := &handles.Clients{}
	parts := &handles.Parts{}
	services := &handles.Services{}
	e.JoinBundle("/stock", roletype.User, mix.Page, clients, parts, services)*/
	r := mux.NewRouter()

	return r
}
