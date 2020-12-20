package mock

import (
	mocket "github.com/Selvatico/go-mocket"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	config := postgres.Config{DriverName: mocket.DriverName}
	dialect := postgres.New(config)

	return gorm.Open(dialect, nil)
}
