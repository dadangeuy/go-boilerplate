package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go-boilerplate/external/postgres"
)

func main() {
	// load environment
	err := godotenv.Load()
	warnOnError(err)

	script := os.Args[1]

	postgresDB, err := postgres.NewPostgresDB()
	panicOnError(err)
	migrator := postgres.NewPostgresMigrator(postgresDB)

	switch script {
	case "migrate":
		err = migrator.Migrate()
		panicOnError(err)

	case "rollback":
		err = migrator.RollbackLast()
		panicOnError(err)

	default:
		err = errors.New(fmt.Sprintf("Script not found: %s", script))
		panic(err)
	}
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
