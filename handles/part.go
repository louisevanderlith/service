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

func GetParts(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Parts", tmpl, "./views/parts.html")

	return func(w http.ResponseWriter, r *http.Request) {

		pagesize := "A10"

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockParts(pagesize)

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

//parts/view/A10
func SearchParts(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Parts", tmpl, "./views/parts.html")

	return func(w http.ResponseWriter, r *http.Request) {

		pagesize := drx.FindParam(r, "pagesize")

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockParts(pagesize)

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

func CreatePart(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("PartCreate", tmpl, "./view/partcreate.html")

	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewPart(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Part View", tmpl, "./view/partview.html")

	return func(w http.ResponseWriter, r *http.Request) {

		key, err := husk.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockPart(key.String())

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
