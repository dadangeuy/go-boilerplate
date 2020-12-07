package component

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"go-boilerplate/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sort"
)

func NewPostgres() (*gorm.DB, error) {
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

func NewPostgresMigrator(postgres *gorm.DB) *gormigrate.Gormigrate {
	migrations := []*gormigrate.Migration{
		migration.NewCreateUserMigration(),
	}
	sort.Slice(migrations, func(i, j int) bool { return migrations[i].ID < migrations[j].ID })

	return gormigrate.New(postgres, gormigrate.DefaultOptions, migrations)
}
