package handles

import (
	"github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/open"
	folio "github.com/louisevanderlith/folio/api"
	"github.com/louisevanderlith/theme/api"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"net/http"
)

var (
	AuthConfig *oauth2.Config
	credConfig *clientcredentials.Config
	Endpoints  map[string]string
)

func SetupRoutes(host, clientId, clientSecret string, endpoints map[string]string) http.Handler {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, endpoints["issuer"])

	if err != nil {
		panic(err)
	}

	Endpoints = endpoints

	AuthConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  host + "/callback",
		Scopes:       []string{oidc.ScopeOpenID, oidc.ScopeOfflineAccess},
	}

	credConfig = &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     provider.Endpoint().TokenURL,
		Scopes:       []string{oidc.ScopeOpenID, "theme", "folio"},
	}

	err = api.UpdateTemplate(credConfig.Client(ctx), endpoints["theme"])

	if err != nil {
		panic(err)
	}

	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	lock := open.NewUILock(provider, AuthConfig)

	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)

	fact := mix.NewPageFactory(tmpl)
	fact.AddMenu(FullMenu())
	fact.AddModifier(mix.EndpointMod(Endpoints))
	fact.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	fact.AddModifier(ThemeContentMod())

	r.Handle("/", lock.Middleware(Index(fact))).Methods(http.MethodGet)

	return r
}

func FullMenu() *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("1", "/stock", "Stock", nil))
	m.AddItem(menu.NewItem("2", "/categories", "Categories", nil))
	m.AddItem(menu.NewItem("3", "/sales", "Sales", nil))

	m.AddItem(menu.NewItem("a", "/regions", "Regions", nil))
	m.AddItem(menu.NewItem("b", "/stock/parts", "Parts", nil))
	m.AddItem(menu.NewItem("b", "/stock/vehicles", "Vehicles", nil))
	m.AddItem(menu.NewItem("b", "/vin", "VIN Numbers", nil))
	m.AddItem(menu.NewItem("e", "/clients", "Clients", nil))

	return m
}

func ThemeContentMod() mix.ModFunc {
	return func(b mix.Bag, r *http.Request) {
		clnt := credConfig.Client(r.Context())

		content, err := folio.FetchDisplay(clnt, Endpoints["folio"])

		if err != nil {
			log.Println("Fetch Profile Error", err)
			panic(err)
			return
		}

		b.SetValue("Folio", content)
	}
}
