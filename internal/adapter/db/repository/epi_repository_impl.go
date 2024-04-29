package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EpiRepositoryImpl struct {
	db *db.ManagerWorker
}

func NewEpiRepositoryImpl(db *db.ManagerWorker) ports.EpiRepository {
	return &EpiRepositoryImpl{db}
}

// GetAll implements ports.EpiRepository.
func (e EpiRepositoryImpl) GetAll(ctx context.Context) ([]*domain.Epi, error) {
	var rst []*domain.Epi

	csr, err := e.db.GetCollection("epi").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = csr.All(ctx, &rst)
	if err != nil {
		return nil, err
	}

	return rst, nil
}

// GetEpi implements ports.EpiRepository.
func (e EpiRepositoryImpl) GetEpi(ctx context.Context, id primitive.ObjectID) (*domain.Epi, error) {
	var epi *domain.Epi

	err := e.db.GetCollection("epi").FindOne(ctx, bson.M{"_id": id}).Decode(&epi)
	if err != nil {
		return nil, err
	}

	return epi, nil
}

// NewEpi implements ports.EpiRepository.
func (e EpiRepositoryImpl) NewEpi(ctx context.Context, epi domain.Epi) (*primitive.ObjectID, error) {

	var oId *primitive.ObjectID
	session, err := e.db.StartSession()
	if err != nil {
		fmt.Printf("59 repo %e\n", err)
		return nil, err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sc mongo.SessionContext) (interface{}, error) {
		epi.Id = primitive.NewObjectID()

		rst, err := sc.Client().Database("america").Collection("epi").InsertOne(ctx, epi)
		if err != nil {
			sc.AbortTransaction(ctx)
			fmt.Printf("69 repo %s\n", err.Error())
			return nil, err
		}

		if err = sc.CommitTransaction(ctx); err != nil {
			fmt.Printf("74 repo %e\n", err)
			return nil, err
		}

		if rstId, ok := rst.InsertedID.(primitive.ObjectID); ok {
			oId = &rstId
			return oId, nil
		} else {
			return nil, errors.New("erro ao converter mongo.insertOneResult em ObjectId")
		}
	})
	if err != nil {
		fmt.Printf("repo %e\n", err)
		return nil, err
	}

	return oId, err
}

// UpdateEpi implements ports.EpiRepository.
func (e EpiRepositoryImpl) UpdateEpi(ctx context.Context, id primitive.ObjectID, epi domain.Epi) (*domain.Epi, error) {
	var session mongo.Session
	var updatedEpi *domain.Epi

	session, err := e.db.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	newData := domain.Epi{
		Id:           id,
		Name:         epi.Name,
		Ca:           epi.Ca,
		EpiType:      epi.EpiType,
		Stock:        epi.Stock,
		MinimumStock: epi.MinimumStock,
		IsReturnable: epi.IsReturnable,
	}

	_, err = session.WithTransaction(ctx, func(sc mongo.SessionContext) (interface{}, error) {
		err = session.Client().Database("america").Collection("epi").FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": &newData}).Decode(&updatedEpi)
		if err != nil {
			fmt.Printf("repository %e", err)
			return nil, err
		}

		if updatedEpi == nil {
			if err = session.AbortTransaction(ctx); err != nil {
				return nil, err
			}
			return nil, errors.New("update operation failed")
		}

		if err = session.CommitTransaction(ctx); err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return &newData, nil
}

// DeleteEpi implements ports.EpiRepository.
func (e EpiRepositoryImpl) DeleteEpi(ctx context.Context, id primitive.ObjectID) error {
	session, err := e.db.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		rst := e.db.GetCollection("epi").FindOneAndDelete(sc, bson.M{"_id": id})
		if rst.Err() != nil {
			session.AbortTransaction(sc)
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	session.EndSession(ctx)

	return nil
}
