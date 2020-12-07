package main

import (
	"go-boilerplate/component"
	"os"
)

func main() {
	command := os.Args[1]

	postgres, err := component.NewPostgres()
	panicOnError(err)
	migrator := component.NewPostgresMigrator(postgres)

	switch command {
	case "migrate":
		err = migrator.Migrate()
		panicOnError(err)

	case "rollback":
		err = migrator.RollbackLast()
		panicOnError(err)
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
