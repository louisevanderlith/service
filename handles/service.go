package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/utility/api"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func GetServices(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagesize := "A10"

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllServices(clnt, Endpoints["stock"], pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, fact.Create(r, "Services", "./views/services.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchServices(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagesize := drx.FindParam(r, "pagesize")
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllServices(clnt, Endpoints["stock"], pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, fact.Create(r, "Services", "./views/services.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreateService(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "Service Create", "./views/servicecreate.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewService(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchService(clnt, Endpoints["stock"], key)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, fact.Create(r, "Service View", "./views/serviceview.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
