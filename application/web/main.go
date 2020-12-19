package main

import (
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"go-boilerplate/component"
	"log"
	"net/http"
)

func main() {
	// load environment
	err := godotenv.Load()
	warnOnError(err)

	// build component
	postgresDB, err := component.NewPostgresDB()
	panicOnError(err)
	infoDelivery := component.NewInfoDelivery()
	userDelivery := component.NewUserDelivery(postgresDB)

	// build router
	router := httprouter.New()
	router.GET("/", infoDelivery.WelcomeHandler)
	router.GET("/users", userDelivery.ListHandler)

	// build server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// run server
	err = server.ListenAndServe()
	panicOnError(err)
}

func warnOnError(err error)  {
	if err != nil {
		log.Print(err)
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
