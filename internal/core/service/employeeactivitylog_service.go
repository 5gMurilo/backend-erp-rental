package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeActivityLogService struct {
	repo     ports.EmployeeActivityLogRepository
	emplRepo ports.EmployeeRepository
}

func NewEmployeeActivityLogService(repo ports.EmployeeActivityLogRepository, emplRepo ports.EmployeeRepository) ports.EmployeeActivityLogService {
	return &EmployeeActivityLogService{repo, emplRepo}
}

func (e EmployeeActivityLogService) GetByEmployee(ctx context.Context, id string) ([]*domain.EmployeeActivityLog, error) {
	employeeId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("24 Log - Service - error: %s\n", err.Error())
		return nil, err
	}

	emp, err := e.emplRepo.GetById(ctx, employeeId)
	if err != nil {
		fmt.Printf("30 Log - Service - error: %s\n", err.Error())
		return nil, err
	}

	activities, err := e.repo.GetByEmployee(ctx, *emp)
	fmt.Println(activities)
	if err != nil {
		fmt.Printf("36 Log - Service - error: %s\n", err.Error())
		return nil, err
	}

	return activities, nil
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
