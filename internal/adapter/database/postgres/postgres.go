package postgres

import (
	"example/tm/authservice/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func New(config config.PostgresConfig) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.PostgresHost,
		config.DatabaseUser,
		config.DatabasePassword,
		config.PostgresDb,
		config.PostgresPort,
	)

	PostgresDb, dbOpenErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbOpenErr != nil {
		return nil, dbOpenErr
	}
	return &Postgres{
		DB: PostgresDb,
	}, nil
}
