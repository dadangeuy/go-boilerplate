package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"reflect"
	"sort"
)

type Migrations []*gormigrate.Migration

func NewMigrations() Migrations {
	migrations := Migrations{}

	migrationsValue := reflect.ValueOf(migrations)
	for methodID := 0; methodID < migrationsValue.NumMethod(); methodID++ {
		migrationMethod := migrationsValue.Method(methodID)
		migration := migrationMethod.
			Call([]reflect.Value{})[0].
			Interface().
			(*gormigrate.Migration)
		migrations = append(migrations, migration)
	}

	sortByIDAsc := func(i, j int) bool { return migrations[i].ID < migrations[j].ID }
	sort.Slice(migrations, sortByIDAsc)

	return migrations
}
