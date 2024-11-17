package repository

import (
	"example/tm/authservice/internal/core/domain"
	"gorm.io/gorm"
	"log/slog"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	err := ur.db.Create(&user)
	if err != nil {
		slog.Info("Error on creating user", err.Error)
		return nil, err.Error
	}
	return user, nil
}
