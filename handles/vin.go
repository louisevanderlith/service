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

func GetVIN(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchLatestVINs(clnt, Endpoints["vin"], "A10")

		if err != nil {
			log.Println("Fetch VINs", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "VIN", "./views/vin/vin.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchVIN(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchLatestVINs(clnt, Endpoints["vin"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch VINs Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "VIN", "./views/vin/vin.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewVIN(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchVIN(clnt, Endpoints["vin"], key)

		if err != nil {
			log.Println("Fetch VIN Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "VIN View", "./views/vinview.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
