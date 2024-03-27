package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/port"
	"context"
	"fmt"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	user, err := s.repo.GetUser(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, error(fmt.Errorf("`%s` already exists", user.Id))
	}

	newUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
