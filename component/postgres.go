package component

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"go-boilerplate/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewPostgresDB() (*gorm.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	name := os.Getenv("POSTGRES_NAME")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, port, name, user, pass,
	)

	return gorm.Open(postgres.Open(dsn), nil)
}

func NewPostgresMigrator(postgresDB *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(
		postgresDB.Debug(),
		gormigrate.DefaultOptions,
		migration.NewMigrations(),
	)
}
