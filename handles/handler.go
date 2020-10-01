package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong/middle"
	"net/http"
)

func SetupRoutes(clnt, scrt, securityUrl, authorityUrl string) http.Handler {
	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))
	clntIns := middle.NewClientInspector(clnt, scrt, http.DefaultClient, securityUrl, authorityUrl)
	r.HandleFunc("/", clntIns.Middleware(Index(tmpl), map[string]bool{"entity.info.search": true, "vehicle.info.search": true})).Methods(http.MethodGet)

	return r
}
