package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
)

type EpiExitsService struct {
	repo ports.EpiExitsRepository
}

func NewEpiExitsService(repo ports.EpiExitsRepository) ports.EpiExitsService {
	return &EpiExitsService{repo}
}

// GetExits implements ports.EpiExitsService.
func (e *EpiExitsService) GetExits(ctx context.Context) ([]*domain.EpiExits, error) {
	rst, err := e.repo.GetExits(ctx)
	if err != nil {
		return nil, err
	}

	return rst, nil
}

// NewExit implements ports.EpiExitsService.
func (e *EpiExitsService) NewExit(ctx context.Context, exit domain.EpiExits) (*domain.EpiExits, error) {
	rst, err := e.repo.NewExit(ctx, exit)
	if err != nil {
		return nil, err
	}

	return rst, nil
}
