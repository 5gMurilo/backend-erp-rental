package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContractRepository interface {
	GetContractById(id primitive.ObjectID) (*domain.ContractData, error)
	GetContractByEmployee(employee domain.Employee) (*domain.ContractData, error)
	GetContractByStatus(status domain.ContractStatus) (*[]domain.ContractData, error)
	GetContracts() ([]domain.ContractData, error)
	CreateContract(contract domain.ContractData) (*domain.ContractData, error)
	UpdateContract(contract domain.ContractData, id primitive.ObjectID) (*domain.ContractData, error)
	DeleteContract(id primitive.ObjectID) error

	GetAssignedContractById(id primitive.ObjectID) (*domain.Contract, error)
	AttachContractToEmployee(id primitive.ObjectID, employee domain.Employee) (*domain.Contract, error)
	ConfirmContractSign(id primitive.ObjectID) (*domain.Contract, error)
}

type ContractService interface {
	GetContracts(ctx context.Context) ([]domain.Contract, error)
	GetContractByEmployee(employee domain.Employee) (*domain.ContractData, error)
	GetContractByStatus(status domain.ContractStatus) (*[]domain.ContractData, error)
	CreateContract(contract domain.Contract) (*domain.Contract, error)
	UpdateContract(contract domain.Contract, id primitive.ObjectID) (*domain.Contract, error)
	DeleteContract(id primitive.ObjectID) error
}
