package main

import (
	"flag"
	"github.com/louisevanderlith/service/handles"
	"net/http"
	"time"
)

func main() {
	host := flag.String("host", "http://127.0.0.1:8096", "This application's URL")
	clientId := flag.String("client", "mango.service", "Client ID which will be used to verify this instance")
	clientSecrt := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	theme := flag.String("theme", "http://127.0.0.1:8093", "Theme URL")
	stock := flag.String("stock", "http://127.0.0.1:8101", "Stock URL")
	flag.Parse()

	ends := map[string]string{
		"issuer": *issuer,
		"theme":  *theme,
		"stock":  *stock,
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8096",
		Handler:      handles.SetupRoutes(*host, *clientId, *clientSecrt, ends),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
