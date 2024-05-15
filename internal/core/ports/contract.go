package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContractRepository interface {
	GetContractByEmployee(employee domain.Employee) (*domain.ContractData, error)
	GetContractByStatus(status domain.ContractStatus) (*[]domain.ContractData, error)
	GetContractById(id primitive.ObjectID) (*domain.Contract, error)
	GetContracts() ([]domain.Contract, error)
	CreateContract(contract domain.Contract) (*mongo.InsertOneResult, error)
	UpdateContract(contract domain.Contract, id primitive.ObjectID) (*domain.Contract, error)
	DeleteContract(id primitive.ObjectID) error
}

type ContractService interface {
	GetContracts(ctx context.Context) ([]domain.Contract, error)
	GetContractByEmployee(employee domain.Employee) (*domain.ContractData, error)
	GetContractByStatus(status domain.ContractStatus) (*[]domain.ContractData, error)
	CreateContract(contract domain.Contract) (*domain.Contract, error)
	UpdateContract(contract domain.Contract, id primitive.ObjectID) (*domain.Contract, error)
	DeleteContract(id primitive.ObjectID) error
}
