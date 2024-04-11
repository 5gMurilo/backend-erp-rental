package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeActivityLogRepository interface {
	GetById(ctx context.Context, id primitive.ObjectID) (*domain.EmployeeActivityLog, error)
	GetByEmployee(ctx context.Context, employee domain.Employee) ([]*domain.EmployeeActivityLog, error)
	New(ctx context.Context, activity domain.EmployeeActivityLog) (*primitive.ObjectID, error)
}

type EmployeeActivityLogService interface {
	GetByEmployee(ctx context.Context, employee domain.Employee) ([]*domain.EmployeeActivityLog, error)
	New(ctx context.Context, activity domain.EmployeeActivityLog) (*domain.EmployeeActivityLog, error)
}
