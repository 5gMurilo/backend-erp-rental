package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContractService struct {
	repo ports.ContractRepository
}

func (c ContractService) GetContracts() ([]domain.ContractData, error) {
	contracts, err := c.repo.GetContracts()
	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func NewContractService(repo ports.ContractRepository) ports.ContractService {
	return &ContractService{repo: repo}
}

func (c ContractService) GetContractByEmployee(employee domain.Employee) ([]domain.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContractService) GetContractByStatus(status domain.ContractStatus) ([]domain.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContractService) CreateContract(contract domain.ContractData) (*domain.ContractData, error) {
	newContract, err := c.repo.CreateContract(contract)
	if err != nil {
		return nil, err
	}

	return newContract, nil
}

func (c ContractService) AttachContractToEmployee(id primitive.ObjectID, employee domain.Employee) (*domain.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContractService) UpdateContract(contract domain.ContractData, id primitive.ObjectID) (*domain.ContractData, error) {
	updatedContract, err := c.repo.UpdateContract(contract, id)
	if err != nil {
		return nil, err
	}

	return updatedContract, nil
}

func (c ContractService) DeleteContract(id primitive.ObjectID) error {
	err := c.repo.DeleteContract(id)
	if err != nil {
		return err
	}

	return nil
}
