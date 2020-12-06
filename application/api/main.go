package main

import (
	"github.com/julienschmidt/httprouter"
	"go-boilerplate/config"
	"go-boilerplate/domain/info"
	"net/http"
)

func main() {
	postgres, err := config.NewPostgres()
	if err != nil {
		panic(err)
	}

	infoDelivery := info.NewDelivery(postgres)

	router := httprouter.New()
	router.GET("/", infoDelivery.WelcomeHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	_ = server.ListenAndServe()
}
