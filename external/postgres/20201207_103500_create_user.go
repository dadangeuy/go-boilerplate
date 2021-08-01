package postgres

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"go-boilerplate/model"
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
