package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"america-rental-backend/internal/core/util"
	"context"
)

type AuthService struct {
	token ports.TokenService
	repo  ports.UserRepository
}

func NewAuthService(repo ports.UserRepository, token ports.TokenService) *AuthService {
	return &AuthService{
		repo:  repo,
		token: token,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (*domain.AuthResponse, error) {
	user, err := as.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = util.ComparePasswords(password, user.Password)
	if err != nil {
		return nil, err
	}

	token, err := as.token.CreateToken(user)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Token:    token,
		UserType: user.UserType,
	}, nil
}
