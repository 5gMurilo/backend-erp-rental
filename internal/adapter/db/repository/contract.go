package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type ContractRepositoryImpl struct {
	db *db.ManagerWorker
}

func NewContractRepository(db *db.ManagerWorker) ports.ContractRepository {
	return &ContractRepositoryImpl{db}
}

// CreateContract implements ports.ContractRepository.
func (c *ContractRepositoryImpl) CreateContract(contract domain.Contract) (*mongo.InsertOneResult, error) {
	session, err := c.db.StartSession()
	if err != nil {
		return nil, err
	}

	coll := session.Client().Database(c.db.Database).Collection("contract")

	err = session.StartTransaction()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	rst, err := coll.InsertOne(ctx, contract)
	if err != nil {
		log.Println(err.Error())
		err := session.AbortTransaction(ctx)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		return nil, err
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return nil, err
	}

	return rst, nil
}

// DeleteContract implements ports.ContractRepository.
func (c *ContractRepositoryImpl) DeleteContract(id primitive.ObjectID) error {
	panic("unimplemented")
}

// GetContractByEmployee implements ports.ContractRepository.
func (c *ContractRepositoryImpl) GetContractByEmployee(employee domain.Employee) (*domain.ContractData, error) {
	panic("unimplemented")
}

// GetContractByStatus implements ports.ContractRepository.
func (c *ContractRepositoryImpl) GetContractByStatus(status domain.ContractStatus) (*[]domain.ContractData, error) {
	panic("unimplemented")
}

// GetContracts implements ports.ContractRepository.
func (c *ContractRepositoryImpl) GetContracts() ([]domain.Contract, error) {
	panic("unimplemented")
}

func (c *ContractRepositoryImpl) GetContractById(id primitive.ObjectID) (*domain.Contract, error) {
	var rst domain.Contract

	err := c.db.GetCollection("contract").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&rst)
	if err != nil {
		return nil, err
	}

	return &rst, nil
}

// UpdateContract implements ports.ContractRepository.
func (c *ContractRepositoryImpl) UpdateContract(contract domain.Contract, id primitive.ObjectID) (*domain.Contract, error) {
	ctx := context.Background()
	session, err := c.db.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	coll := session.Client().Database(c.db.Database).Collection("contract")

	err = session.StartTransaction()
	if err != nil {
		return nil, err
	}

	sr := coll.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": contract})

	if sr.Err() != nil {
		log.Println(err.Error())
		err := session.AbortTransaction(ctx)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		return nil, err
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return nil, err
	}

	var newData domain.Contract

	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(&newData)
	if err != nil {
		return nil, err
	}

	return &newData, nil
}
