package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/vehicle/api"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func GetVehicles(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllVehicles(clnt, Endpoints["vehicle"], "A10")

		if err != nil {
			log.Println("Fetch Vehicles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Vehicles", "./views/vehicles.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchVehicles(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllVehicles(clnt, Endpoints["vehicle"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Vehicles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Vehicles", "./views/vehicles.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewVehicles(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchVehicleInfo(clnt, Endpoints["vehicle"], key)

		if err != nil {
			log.Println("Fetch Vehicle Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Vehicles View", "./views/vehicleview.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
