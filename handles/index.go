package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Index", tmpl, "./views/index.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
