package main

import (
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"go-boilerplate/domain/info"
	"go-boilerplate/domain/user"
	"go-boilerplate/external/database"
	"log"
	"net/http"
)

func main() {
	// load environment
	err := godotenv.Load()
	warnOnError(err)

	// build external component
	postgresDB, err := database.NewDB()
	panicOnError(err)

	// build info component
	infoDelivery := info.NewInfoDelivery()

	// build user component
	userRepository := user.NewUserRepository(postgresDB)
	userUsecase := user.NewUserUsecase(userRepository)
	userDelivery := user.NewUserDelivery(userUsecase)

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
