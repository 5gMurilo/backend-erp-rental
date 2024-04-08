package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"
)

type TokenService interface {
	CreateToken(user *domain.User) (string, error)
	VerifyToken(token string) (*domain.TokenPayload, error)
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (*domain.AuthResponse, error)
}
