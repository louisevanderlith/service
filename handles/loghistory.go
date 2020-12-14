package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func GetHistory(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "Index", "./views/index.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
