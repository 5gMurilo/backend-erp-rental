package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnableEpiRepository interface {
	CreateReturnableEpi(ctx context.Context, returnableEpi domain.ReturnableEpi) (*primitive.ObjectID, error)
	UpdateReturnableEpi(ctx context.Context, returnableEpi domain.ReturnableEpi, id primitive.ObjectID) error
	GetReturnableEpi(ctx context.Context, id primitive.ObjectID) (*domain.ReturnableEpi, error)
	GetAllReturnableEpi(ctx context.Context) ([]*domain.ReturnableEpi, error)
	DeleteReturnableEpi(ctx context.Context, id primitive.ObjectID) error
}

type ReturnableEpiService interface {
	CreateReturnableEpi(ctx context.Context, returnableEpi domain.ReturnableEpi) (*domain.ReturnableEpi, error)
	UpdateReturnableEpi(ctx context.Context, returnableEpi domain.ReturnableEpi, id primitive.ObjectID) (*domain.ReturnableEpi, error)
	GetReturnableEpi(ctx context.Context, id primitive.ObjectID) (*domain.ReturnableEpi, error)
	GetAllReturnableEpi(ctx context.Context) ([]*domain.ReturnableEpi, error)
	DeleteReturnableEpi(ctx context.Context, id primitive.ObjectID) error
}
