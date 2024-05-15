package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StorageRepository struct {
	db *db.ManagerWorker
}

const _collection = "storage"

func NewStorageRepository(db *db.ManagerWorker) ports.StorageRepository {
	return &StorageRepository{db}
}

func (s StorageRepository) RegisterUpdateInformation(ctx context.Context, onedriveFile domain.OnedriveFile, actor string) (*domain.OnedriveFile, error) {
	session, err := s.db.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	onedriveFile.Id = primitive.NewObjectID()
	onedriveFile.UpdatedBy = actor

	_, err = session.WithTransaction(context.TODO(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		rst, err := sessionContext.Client().Database("america").Collection(_collection).InsertOne(context.TODO(), onedriveFile)
		if err != nil {
			fmt.Println(err)
			err := session.AbortTransaction(sessionContext)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return nil, err
		}

		err = session.CommitTransaction(sessionContext)
		if err != nil {
			return nil, err
		}

		return rst.InsertedID, nil
	})

	if err != nil {
		return nil, err
	}

	return &onedriveFile, nil
}

func (s StorageRepository) GetOnedriveFilesByEmployee(ctx context.Context, employeeName string) (*[]domain.OnedriveFile, error) {
	cursor, err := s.db.GetCollection(_collection).Find(ctx, bson.M{"employee": employeeName})
	if err != nil {
		return nil, err
	}
	var results []domain.OnedriveFile
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (s StorageRepository) UpdateOnedriveFile(ctx context.Context, file domain.OnedriveFile, _ string) (*domain.OnedriveFile, error) {
	session, err := s.db.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	var result domain.OnedriveFile
	_, err = session.WithTransaction(context.TODO(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		err := sessionContext.Client().Database("america").Collection(_collection).FindOneAndReplace(context.TODO(), bson.M{"_id": file.Id}, file).Decode(&result)
		if err != nil {
			return nil, err
		}

		return nil, err
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s StorageRepository) DeleteOnedriveFile(ctx context.Context, driveItemid string) error {
	sr := s.db.GetCollection(_collection).FindOneAndDelete(ctx, bson.M{"driveItemId": driveItemid})
	if sr.Err() != nil {
		return sr.Err()
	}

	return nil
}
