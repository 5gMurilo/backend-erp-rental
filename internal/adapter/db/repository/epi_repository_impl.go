package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
