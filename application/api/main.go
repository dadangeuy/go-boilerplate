package main

import (
	"github.com/julienschmidt/httprouter"
	"go-boilerplate/domain/info"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: buildHandler(),
	}

	_ = server.ListenAndServe()
}

func buildHandler() http.Handler {
	infoDelivery := info.NewDelivery()

	router := httprouter.New()
	router.GET("/", infoDelivery.WelcomeHandler)

	return router
}
