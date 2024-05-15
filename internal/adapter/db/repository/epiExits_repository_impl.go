package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type EpiExitsRepositoryImpl struct {
	db *db.ManagerWorker
}

func NewEpiExitsRepositoryImpl(db *db.ManagerWorker) ports.EpiExitsRepository {
	return &EpiExitsRepositoryImpl{db}
}

// getExits implements ports.EpiExitsRepository.
func (e EpiExitsRepositoryImpl) GetExits(ctx context.Context) ([]*domain.EpiExits, error) {
	var exits []*domain.EpiExits

	csr, err := e.db.GetCollection("epiExits").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = csr.All(ctx, &exits)
	if err != nil {
		return nil, err
	}

	return exits, nil
}

func (e EpiExitsRepositoryImpl) GetExitById(ctx context.Context, id primitive.ObjectID) (*domain.EpiExits, error) {
	var exit *domain.EpiExits

	err := e.db.GetCollection("epiExits").FindOne(ctx, bson.M{"_id": id}).Decode(&exit)
	if err != nil {
		return nil, err
	}

	return exit, nil
}

// newExit implements ports.EpiExitsRepository.
func (e EpiExitsRepositoryImpl) NewExit(ctx context.Context, exit domain.EpiExits) (*primitive.ObjectID, error) {
	session, err := e.db.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	exit.ExitTime = primitive.NewDateTimeFromTime(time.Now())

	err = session.StartTransaction()
	if err != nil {
		return nil, err
	}

	result, err := session.Client().Database("america").Collection("epiExits").InsertOne(ctx, exit)
	if err != nil {
		err = session.AbortTransaction(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	if result == nil {
		err = session.AbortTransaction(ctx)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("unknown error during insert operation")
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return &id, nil
}
