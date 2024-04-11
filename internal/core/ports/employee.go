package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeRepository interface {
	New(employee domain.Employee) (primitive.ObjectID, error)
	GetAll(ctx context.Context) ([]*domain.Employee, error)
	GetById(ctx context.Context, id primitive.ObjectID) (*domain.Employee, error)
	GetByCPF(ctx context.Context, cpf string) (*domain.Employee, error)
	Update(ctx context.Context, id primitive.ObjectID, newData domain.Employee) (*domain.Employee, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type EmployeeService interface {
	New(employee domain.Employee) (primitive.ObjectID, error)
	GetAll(ctx context.Context) ([]*domain.Employee, error)
	GetById(ctx context.Context, id primitive.ObjectID) (*domain.Employee, error)
	GetByCPF(ctx context.Context, cpf string) (*domain.Employee, error)
	Update(ctx context.Context, id primitive.ObjectID, newData domain.Employee) (*domain.Employee, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}
