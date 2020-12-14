package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/parts/api"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func GetParts(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pagesize := "A10"

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllSpares(clnt, Endpoints["stock"], pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, fact.Create(r, "Parts", "./views/parts.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

//parts/view/A10
func SearchParts(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagesize := drx.FindParam(r, "pagesize")
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllSpares(clnt, Endpoints["stock"], pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, fact.Create(r, "Parts", "./views/parts.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreatePart(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "PartCreate", "./view/partcreate.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewPart(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchSpare(clnt, Endpoints["stock"], key)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, fact.Create(r, "Part View", "./view/partview.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
