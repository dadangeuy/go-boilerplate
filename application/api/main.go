package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	"go-boilerplate/domain/info"
	"go-boilerplate/domain/user"
	"go-boilerplate/external/postgres"
)

func main() {
	// load environment
	err := godotenv.Load()
	warnOnError(err)

	// build component
	postgresDB, err := postgres.NewPostgresDB()
	panicOnError(err)
	infoDelivery := info.NewInfoDelivery()
	userDelivery := user.NewUserDelivery(postgresDB)

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

func warnOnError(err error) {
	if err != nil {
		log.Print(err)
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
