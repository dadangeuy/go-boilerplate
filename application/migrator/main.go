package main

import (
	"github.com/joho/godotenv"
	"go-boilerplate/external/database"
	"log"
	"os"
)

func main() {
	// load environment
	err := godotenv.Load()
	warnOnError(err)

	command := os.Args[1]

	postgresDB, err := database.NewDB()
	panicOnError(err)
	migrator := database.NewMigrator(postgresDB)

	switch command {
	case "migrate":
		err = migrator.Migrate()
		panicOnError(err)

	case "rollback":
		err = migrator.RollbackLast()
		panicOnError(err)
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
