package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeService struct {
	repo ports.EmployeeRepository
}

func NewEmployeeService(repo ports.EmployeeRepository) ports.EmployeeService {
	return &EmployeeService{repo}
}

// Delete implements ports.EmployeeService.
func (e EmployeeService) Delete(ctx context.Context, id primitive.ObjectID) error {
	err := e.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements ports.EmployeeService.
func (e EmployeeService) GetAll(ctx context.Context) ([]*domain.Employee, error) {
	employees, err := e.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

// GetByCPF implements ports.EmployeeService.
func (e EmployeeService) GetByCPF(ctx context.Context, cpf string) (*domain.Employee, error) {
	employee, err := e.repo.GetByCPF(ctx, cpf)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

// GetById implements ports.EmployeeService.
func (e EmployeeService) GetById(ctx context.Context, id primitive.ObjectID) (*domain.Employee, error) {
	employee, err := e.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

// New implements ports.EmployeeService.
func (e EmployeeService) New(ctx context.Context, employee domain.Employee, createdBy string) (*domain.Employee, error) {
	employee.ModifiedBy = createdBy

	id, err := e.repo.New(ctx, employee)
	if err != nil {
		fmt.Printf("service error 66 \n%s\n", err.Error())
		return nil, err
	}

	pId, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		fmt.Printf("service error 72 \n%s\n", err.Error())
		return nil, err
	}

	log.Println(pId)

	emp, err := e.repo.GetById(ctx, pId)
	if err != nil {
		fmt.Printf("service error 78\n%s\n", err.Error())
		return nil, err
	}

	return emp, nil
}

// Update implements ports.EmployeeService.
func (e EmployeeService) Update(ctx context.Context, id primitive.ObjectID, data domain.Employee, updatedBy string) (*domain.Employee, error) {
	data.ModifiedBy = updatedBy

	newEmpData, err := e.repo.Update(ctx, id, data)
	if err != nil {
		return nil, err
	}

	return newEmpData, err
}
