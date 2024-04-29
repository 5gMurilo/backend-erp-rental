package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
	worker *db.ManagerWorker
}

func NewUserRepositoryImpl(worker *db.ManagerWorker) ports.UserRepository {
	return &UserRepositoryImpl{
		worker: worker,
	}
}

func (u UserRepositoryImpl) Get(c context.Context, id primitive.ObjectID) (*domain.User, error) {
	var result *domain.User
	err := u.worker.GetCollection("user").FindOne(c, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u UserRepositoryImpl) GetByEmail(c context.Context, email string) (*domain.User, error) {
	var result *domain.User
	err := u.worker.GetCollection("user").FindOne(c, bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u UserRepositoryImpl) GetAll(c context.Context) (*[]domain.User, error) {
	var rst []domain.User

	cursor, err := u.worker.GetCollection("user").Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	err = cursor.All(c, &rst)
	if err != nil {
		return nil, err
	}

	return &rst, nil
}

func (u UserRepositoryImpl) Create(c context.Context, data domain.User) (*primitive.ObjectID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newUser := domain.User{
		Id:        primitive.NewObjectID(),
		Name:      data.Name,
		Email:     data.Email,
		Password:  string(hash),
		State:     data.State,
		Role:      data.Role,
		UserType:  data.UserType,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		//change updatedBy for name from decoded Token
		UpdatedBy: "Me",
	}

	rst, err := u.worker.GetCollection("user").InsertOne(c, newUser)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if oId, ok := rst.InsertedID.(primitive.ObjectID); ok {
		return &oId, nil
	} else {
		return nil, errors.New("erro ao decodificar o Id")
	}
}

func (u UserRepositoryImpl) Update(c context.Context, data domain.User, id primitive.ObjectID) (*domain.User, error) {
	newData := domain.User{
		Id:        id,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		State:     data.State,
		Role:      data.Role,
		UserType:  data.UserType,
		CreatedAt: data.CreatedAt,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedBy: "me",
	}

	rst, err := u.worker.GetCollection("user").UpdateOne(c, bson.M{"_id": id}, bson.M{"$set": newData})
	if err != nil {
		fmt.Printf("repository %e", err)
		return nil, err
	}

	if rst.MatchedCount == 0 {
		return nil, errors.New("update operation failed")
	}

	return &newData, nil
}

func (u UserRepositoryImpl) Delete(c context.Context, id primitive.ObjectID) error {
	_, err := u.worker.GetCollection("user").DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
