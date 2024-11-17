package domain

import "time"

type UserRole string

const (
	Admin  UserRole = "admin"
	Client UserRole = "client"
)

type User struct {
	ID          uint64
	Email       string
	Password    string
	UserRole    UserRole
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
