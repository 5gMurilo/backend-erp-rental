package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type ContractRepositoryImpl struct {
	db db.ManagerWorker
}

func NewContractRepository(db db.ManagerWorker) ports.ContractRepository {
	return &ContractRepositoryImpl{db}
}

func (c ContractRepositoryImpl) GetContractById(id primitive.ObjectID) (*domain.ContractData, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContractRepositoryImpl) GetContractByEmployee(employee domain.Employee) (*domain.ContractData, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContractRepositoryImpl) GetContractByStatus(status domain.ContractStatus) (*[]domain.ContractData, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContractRepositoryImpl) GetContracts() ([]domain.ContractData, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContractRepositoryImpl) CreateContract(contract domain.ContractData) (*domain.ContractData, error) {
	session, err := c.db.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())
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

	id, err := primitive.ObjectIDFromHex(rst.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}

	var value domain.ContractData
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(&value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (c ContractRepositoryImpl) UpdateContract(contract domain.ContractData, id primitive.ObjectID) (*domain.ContractData, error) {
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

	var newData domain.ContractData
	err = sr.Decode(&newData)
	if err != nil {
		return nil, err
	}

	return &newData, nil
}

func (c ContractRepositoryImpl) DeleteContract(id primitive.ObjectID) error {
	session, err := c.db.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())
	coll := session.Client().Database(c.db.Database).Collection("contract")
	err = session.StartTransaction()
	if err != nil {
		return err
	}

	result := coll.FindOneAndDelete(context.TODO(), bson.M{"_id": id})
	if result.Err() != nil {
		log.Println(result.Err())
		err := session.AbortTransaction(context.TODO())
		if err != nil {
			log.Println(result.Err())
			return err
		}
		return result.Err()
	}
	err = session.CommitTransaction(context.TODO())
	if err != nil {
		return err
	}
	
	return nil
}

func (c ContractRepositoryImpl) GetAssignedContractById(id primitive.ObjectID) (*domain.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContractRepositoryImpl) AttachContractToEmployee(id primitive.ObjectID, employee domain.Employee) (*domain.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContractRepositoryImpl) ConfirmContractSign(id primitive.ObjectID) (*domain.Contract, error) {
	//TODO implement me
	panic("implement me")
}
