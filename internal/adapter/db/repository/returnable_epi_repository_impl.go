package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReturnableEpiRepository struct {
	db db.ManagerWorker
}

func NewReturnableEpiRepositoryImpl(db db.ManagerWorker) ports.ReturnableEpiRepository {
	return ReturnableEpiRepository{db}
}

// CreateReturnableEpi implements ports.ReturnableEpiRepository.
func (r ReturnableEpiRepository) CreateReturnableEpi(ctx context.Context, returnableEpi domain.ReturnableEpi) (*primitive.ObjectID, error) {
	session, err := r.db.StartSession()
	if err != nil {
		return nil, err
	}

	objId := primitive.NewObjectID()

	rst, err := session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		returnableEpi.ID = objId

		rst, err := session.Client().Database("america").Collection("returnableEpi").InsertOne(ctx, returnableEpi)
		if err != nil {
			return nil, err
		}

		return rst.InsertedID, nil
	})
	if err != nil {
		return nil, err
	}

	if oId, ok := rst.(primitive.ObjectID); ok {
		return &oId, nil
	} else {
		return nil, errors.New("erro ao decodificar o Id")
	}
}

// DeleteReturnableEpi implements ports.ReturnableEpiRepository.
func (r ReturnableEpiRepository) DeleteReturnableEpi(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.db.GetCollection("returnableEpi").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return err
}

// GetAllReturnableEpi implements ports.ReturnableEpiRepository.
func (r ReturnableEpiRepository) GetAllReturnableEpi(ctx context.Context) (*[]domain.ReturnableEpi, error) {
	session, err := r.db.StartSession()
	if err != nil {
		return nil, err
	}

	var data []domain.ReturnableEpi

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		cursor, err := session.Client().Database("america").Collection("returnableEpi").Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}

		err = cursor.All(ctx, &data)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetReturnableEpi implements ports.ReturnableEpiRepository.
func (r ReturnableEpiRepository) GetReturnableEpi(ctx context.Context, id primitive.ObjectID) (*domain.ReturnableEpi, error) {
	var returnableEpi domain.ReturnableEpi
	err := r.db.GetCollection("returnableEpi").FindOne(ctx, bson.M{"_id": id}).Decode(&returnableEpi)
	if err != nil {
		return nil, err
	}

	return &returnableEpi, nil
}

// UpdateReturnableEpi implements ports.ReturnableEpiRepository.
func (r ReturnableEpiRepository) UpdateReturnableEpi(ctx context.Context, returnableEpi domain.ReturnableEpi, id primitive.ObjectID) (*domain.ReturnableEpi, error) {
	var mRst domain.ReturnableEpi
	session, err := r.db.StartSession()
	if err != nil {
		return nil, err
	}

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		err := session.Client().Database("america").Collection("returnableEpi").FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": returnableEpi}).Decode(mRst)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}
	return &mRst, nil
}
