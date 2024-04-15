package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EpiRepository interface {
	GetAll(ctx context.Context) ([]*domain.Epi, error)
	GetEpi(ctx context.Context, id primitive.ObjectID) (*domain.Epi, error)
	NewEpi(ctx context.Context, epi domain.Epi) (*primitive.ObjectID, error)
	UpdateEpi(ctx context.Context, id primitive.ObjectID, epi domain.Epi) (*domain.Epi, error)
	DeleteEpi(ctx context.Context, id primitive.ObjectID) error
}

type EpiService interface {
	GetAll(ctx context.Context) ([]*domain.Epi, error)
	GetEpi(ctx context.Context, id primitive.ObjectID) (*domain.Epi, error)
	NewEpi(ctx context.Context, epi domain.Epi) (*primitive.ObjectID, error)
	UpdateEpi(ctx context.Context, id primitive.ObjectID, epi domain.Epi) (*domain.Epi, error)
	DeleteEpi(ctx context.Context, id primitive.ObjectID) error
}
