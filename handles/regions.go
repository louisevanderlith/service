package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/vin/api"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func GetRegions(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchRegions(clnt, Endpoints["vin"], "A10")

		if err != nil {
			log.Println("Fetch Regions Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Regions", "./views/regions.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchRegions(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchRegions(clnt, Endpoints["vin"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Regions", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Regions", "./views/regions.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewRegion(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchRegion(clnt, Endpoints["vin"], key)

		if err != nil {
			log.Println("Fetch Region Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Region View", "./views/regionview.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
