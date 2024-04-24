package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnableEpiService struct {
	repo ports.ReturnableEpiRepository
}

func NewReturnableEpiService(repo ports.ReturnableEpiRepository) ports.ReturnableEpiService {
	return &ReturnableEpiService{repo}
}

// CreateReturnableEpi implements ports.ReturnableEpiService.
func (r *ReturnableEpiService) CreateReturnableEpi(ctx context.Context, returnableEpi domain.ReturnableEpi) (*domain.ReturnableEpi, error) {
	rst, err := r.repo.CreateReturnableEpi(ctx, returnableEpi)
	if err != nil {
		return nil, err
	}

	rEpi, err := r.repo.GetReturnableEpi(ctx, *rst)
	if err != nil {
		return nil, err
	}

	return rEpi, err
}

// DeleteReturnableEpi implements ports.ReturnableEpiService.
func (r *ReturnableEpiService) DeleteReturnableEpi(ctx context.Context, id primitive.ObjectID) error {
	err := r.repo.DeleteReturnableEpi(ctx, id)

	return err
}

// GetAllReturnableEpi implements ports.ReturnableEpiService.
func (r *ReturnableEpiService) GetAllReturnableEpi(ctx context.Context) (*[]domain.ReturnableEpi, error) {
	rst, err := r.repo.GetAllReturnableEpi(ctx)
	if err != nil {
		return nil, err
	}

	return rst, err
}

// GetReturnableEpi implements ports.ReturnableEpiService.
func (r *ReturnableEpiService) GetReturnableEpi(ctx context.Context, id primitive.ObjectID) (*domain.ReturnableEpi, error) {
	epi, err := r.repo.GetReturnableEpi(ctx, id)
	if err != nil {
		return nil, err
	}

	return epi, nil
}

// UpdateReturnableEpi implements ports.ReturnableEpiService.
func (r *ReturnableEpiService) UpdateReturnableEpi(ctx context.Context, returnableEpi domain.ReturnableEpi, id primitive.ObjectID) (*domain.ReturnableEpi, error) {
	rst, err := r.repo.UpdateReturnableEpi(ctx, returnableEpi, id)
	if err != nil {
		return nil, err
	}

	return rst, nil
}
