package service

import "example/tm/authservice/internal/adapter/database/postgres/repository"

type UserService struct {
	resp *repository.UserRepository
}

func NewUserService(resp *repository.UserRepository) *UserService {
	return &UserService{
		resp,
	}
}
