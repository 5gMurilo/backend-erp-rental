package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/port"
	"america-rental-backend/internal/core/util"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	if hashedPassword == "" {
		return nil, error(fmt.Errorf("`%s` is not a valid password", user.Password))
	}

	user.Password = hashedPassword

	newUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *UserService) GetUser(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	// var users []domain.User
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *domain.User, newUserInfo *domain.User) (*domain.User, error) {
	user, err := s.repo.GetUser(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
