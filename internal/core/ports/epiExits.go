package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EpiExitsRepository interface {
	NewExit(ctx context.Context, exit domain.EpiExits) (*primitive.ObjectID, error)
	GetExits(ctx context.Context) ([]*domain.EpiExits, error)
}

type EpiExitsService interface {
	NewExit(ctx context.Context, exit domain.EpiExits) (*domain.EpiExits, error)
	GetExits(ctx context.Context) ([]*domain.EpiExits, error)
}
