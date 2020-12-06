package repository

import "gorm.io/gorm"

type Repository struct {
	Postgres *gorm.DB
}
