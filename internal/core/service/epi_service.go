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
