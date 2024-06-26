package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	exit.ID = primitive.NewObjectID()
	exit.ExitTime = primitive.NewDateTimeFromTime(time.Now())

	_, err = session.WithTransaction(ctx, func(sessionContext mongo.SessionContext) (interface{}, error) {
		rst, err := e.db.GetCollection("epiExits").InsertOne(sessionContext, exit)
		if err != nil {
			session.AbortTransaction(sessionContext)
			return nil, err
		}

		if rst == nil {
			session.AbortTransaction(sessionContext)
			return nil, errors.New("unknown error during insert operation")
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return &exit.ID, nil
}
