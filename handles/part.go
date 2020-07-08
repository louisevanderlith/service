package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/service/resources"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

func GetParts(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Parts",tmpl)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		pagesize := "A10"

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchStockParts(pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

//parts/view/A10
func SearchParts(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Parts",tmpl)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		pagesize := ctx.FindParam("pagesize")

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchStockParts(pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreatePart(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("PartCreate",tmpl)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		err := ctx.Serve(http.StatusOK, pge.Page(nil, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewPart(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Index",tmpl)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchStockPart(key.String())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
