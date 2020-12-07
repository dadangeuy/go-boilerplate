package main

import (
	"github.com/julienschmidt/httprouter"
	"go-boilerplate/component"
	"net/http"
)

func main() {
	// build component
	postgresDB, err := component.NewPostgresDB()
	panicOnError(err)
	infoDelivery := component.NewInfoDelivery(postgresDB)

	// build router
	router := httprouter.New()
	router.GET("/", infoDelivery.WelcomeHandler)

	// build server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// run server
	err = server.ListenAndServe()
	panicOnError(err)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
