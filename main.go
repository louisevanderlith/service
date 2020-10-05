package main

import (
	"flag"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/service/handles"
	"net/http"
	"time"
)

func main() {
	clientId := flag.String("client", "mango.service", "Client ID which will be used to verify this instance")
	clientSecrt := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	security := flag.String("security", "http://localhost:8086", "Security Provider's URL")
	manager := flag.String("manager", "http://localhost:8097", "User Provider's URL")
	authority := flag.String("authority", "http://localhost:8094", "Authority Provider's URL")

	flag.Parse()

	err := drx.UpdateTemplate(*clientId, *clientSecrt, *security)

	if err != nil {
		panic(err)
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8096",
		Handler:      handles.SetupRoutes(*clientId, *clientSecrt, *security, *manager, *authority),
	}

	err = srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
