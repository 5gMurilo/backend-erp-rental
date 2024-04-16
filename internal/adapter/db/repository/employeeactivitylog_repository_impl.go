package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeActivityLog struct {
	db *db.ManagerWorker
}

func NewEmployeeActivityLog(db *db.ManagerWorker) ports.EmployeeActivityLogRepository {
	return &EmployeeActivityLog{db}
}

func (e EmployeeActivityLog) GetByEmployee(ctx context.Context, employee domain.Employee) ([]*domain.EmployeeActivityLog, error) {
	var data []*domain.EmployeeActivityLog

	rst, err := e.db.GetCollection("employeeActivityLog").Find(ctx, bson.M{"employee": employee})
	if err != nil {
		return nil, err
	}
	defer rst.Close(ctx)

	err = rst.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e EmployeeActivityLog) GetById(ctx context.Context, id primitive.ObjectID) (*domain.EmployeeActivityLog, error) {
	var activity domain.EmployeeActivityLog

	err := e.db.GetCollection("employeeActivityLog").FindOne(ctx, bson.M{"_id": id}).Decode(&activity)
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (e EmployeeActivityLog) New(ctx context.Context, activity domain.EmployeeActivityLog) (*primitive.ObjectID, error) {
	rst, err := e.db.GetCollection("employeeActivityLog").InsertOne(ctx, activity)
	if err != nil {
		return nil, err
	}

	if id, ok := rst.InsertedID.(primitive.ObjectID); ok {
		return &id, nil
	} else {
		return nil, errors.New("erro ao decodificar o Id")
	}
}
