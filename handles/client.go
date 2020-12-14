package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"log"
	"net/http"
)

func GetClients(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "Clients", "./views/clients.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchClients(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "Clients", "./views/clients.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewClient(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		//Quotes...
		err = mix.Write(w, fact.Create(r, "ClientEdit", "./view/clientedit.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreateClient(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r,"ClientCreate", "./views/clientcreate.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
