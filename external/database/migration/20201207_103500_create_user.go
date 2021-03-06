package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"go-boilerplate/model"
	"gorm.io/gorm"
)

func (Migrations) CreateUserMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20201207_103500_create_user",
		Migrate: func(db *gorm.DB) error {
			return db.Migrator().CreateTable(&model.User{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&model.User{})
		},
	}
}
