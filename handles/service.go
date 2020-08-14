package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/service/resources"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func GetServices(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("services", tmpl, "./views/services.html")

	return func(w http.ResponseWriter, r *http.Request) {

		pagesize := "A10"

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockServices(pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchServices(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Services", tmpl, "./views/services.html")

	return func(w http.ResponseWriter, r *http.Request) {

		pagesize := drx.FindParam(r, "pagesize")

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockService(pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreateService(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Service Create", tmpl, "./views/servicecreate.html")

	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewService(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Service View", tmpl, "./views/serviceview.html")

	return func(w http.ResponseWriter, r *http.Request) {

		key, err := husk.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockService(key.String())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
