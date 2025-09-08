package ports

import "hexagonal/core/domain"

type UserServies interface {
	CreateUser(user *domain.User) error
	GetUser(id string) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
}
