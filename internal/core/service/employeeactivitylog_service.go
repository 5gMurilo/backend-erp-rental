package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeActivityLogService struct {
	repo ports.EmployeeActivityLogRepository
}

func NewEmployeeActivityLogService(repo ports.EmployeeActivityLogRepository) ports.EmployeeActivityLogService {
	return &EmployeeActivityLogService{repo}
}

func (e EmployeeActivityLogService) GetByEmployee(ctx context.Context, employee domain.Employee) ([]*domain.EmployeeActivityLog, error) {
	activities, err := e.repo.GetByEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}

	return activities, err
}

func (e EmployeeActivityLogService) New(ctx context.Context, activity domain.EmployeeActivityLog) (*domain.EmployeeActivityLog, error) {
	id, err := e.repo.New(ctx, activity)
	if err != nil {
		return nil, err
	}

	pId, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		return nil, err
	}

	createdActivity, err := e.repo.GetById(ctx, pId)
	if err != nil {
		return nil, err
	}

	return createdActivity, nil
}
