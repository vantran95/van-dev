package ports

import "hexagonal/core/domain"

type UserRepository interface {
	Save(user *domain.User) error
	FindByID(id string) (*domain.User, error)
	FindAll() ([]*domain.User, error)
}
