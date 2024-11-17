package postgres

import (
	"example/tm/authservice/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDb *gorm.DB

type Postgres struct {
	DB *gorm.DB
}

func New(config config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s passowrd=%s dbname=%s port=%s",
		config.PostgresHost,
		config.DatabaseUser,
		config.DatabasePassword,
		config.PostgresDb,
		config.PostgresPort,
	)

	PostgresDb, dbOpenErr := gorm.Open(postgres.Open(dsn))

	if dbOpenErr != nil {
		return nil, dbOpenErr
	}
	return PostgresDb, nil
}
