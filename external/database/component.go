package database

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"go-boilerplate/external/database/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"reflect"
	"sort"
)

func NewDB() (*gorm.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	name := os.Getenv("POSTGRES_NAME")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, port, name, user, pass,
	)
	dialector := postgres.Open(dsn)

	return gorm.Open(dialector, nil)
}

func NewMigrator(postgresDB *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(
		postgresDB.Debug(),
		gormigrate.DefaultOptions,
		NewMigrations(),
	)
}

func NewMigrations() migration.Migrations {
	migrations := migration.Migrations{}

	migrationsValue := reflect.ValueOf(migrations)
	for methodID := 0; methodID < migrationsValue.NumMethod(); methodID++ {
		migrationMethod := migrationsValue.Method(methodID)
		migrationValue := migrationMethod.
			Call([]reflect.Value{})[0].
			Interface().(*gormigrate.Migration)
		migrations = append(migrations, migrationValue)
	}

	sortByIDAsc := func(i, j int) bool { return migrations[i].ID < migrations[j].ID }
	sort.Slice(migrations, sortByIDAsc)

	return migrations
}
