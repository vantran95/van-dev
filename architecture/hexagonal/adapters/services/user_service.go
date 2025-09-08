package services

import (
	"hexagonal/core/domain"
	"hexagonal/core/ports"
)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserServies {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) CreateUser(user *domain.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	return s.repo.Save(user)
}

func (s *UserServiceImpl) GetUser(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserServiceImpl) ListUsers() ([]*domain.User, error) {
	return s.repo.FindAll()
}
