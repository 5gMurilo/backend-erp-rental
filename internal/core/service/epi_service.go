package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EpiService struct {
	repo ports.EpiRepository
}

func NewEpiService(repo ports.EpiRepository) ports.EpiService {
	return &EpiService{repo}
}

// DeleteEpi implements ports.EpiService.
func (e *EpiService) DeleteEpi(ctx context.Context, id primitive.ObjectID) error {
	err := e.repo.DeleteEpi(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements ports.EpiService.
func (e *EpiService) GetAll(ctx context.Context) ([]*domain.Epi, error) {
	rst, err := e.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return rst, nil
}

// GetEpi implements ports.EpiService.
func (e *EpiService) GetEpi(ctx context.Context, id primitive.ObjectID) (*domain.Epi, error) {
	epi, err := e.repo.GetEpi(ctx, id)
	if err != nil {
		return nil, err
	}

	return epi, nil
}

// NewEpi implements ports.EpiService.
func (e *EpiService) NewEpi(ctx context.Context, epi domain.Epi) (*domain.Epi, error) {
	rst, err := e.repo.NewEpi(ctx, epi)
	if err != nil {
		return nil, err
	}

	createdEpi, err := e.repo.GetEpi(ctx, *rst)
	if err != nil {
		return nil, err
	}

	return createdEpi, nil
}

// UpdateEpi implements ports.EpiService.
func (e *EpiService) UpdateEpi(ctx context.Context, id primitive.ObjectID, epi domain.Epi) (*domain.Epi, error) {
	rst, err := e.repo.UpdateEpi(ctx, id, epi)
	if err != nil {
		return nil, err
	}

	return rst, nil
}
