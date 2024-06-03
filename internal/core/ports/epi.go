package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EpiRepository interface {
	GetAll(ctx context.Context) ([]*domain.Epi, error)
	GetEpi(ctx context.Context, id primitive.ObjectID) (*domain.Epi, error)
}

type EpiService interface {
	GetAll(ctx context.Context) ([]*domain.Epi, error)
	GetEpi(ctx context.Context, id primitive.ObjectID) (*domain.Epi, error)
}
