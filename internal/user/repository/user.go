package repository

import (
	"america-rental-backend/adapters/db"
	"america-rental-backend/internal/user"
	"america-rental-backend/internal/user/ports"
	"america-rental-backend/internal/util"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRepository struct {
	worker *db.ManagerWorker
}

const collection string = "user"

func NewUserRepository(worker *db.ManagerWorker) ports.UserRepository {
	return &UserRepository{
		worker: worker,
	}
}

func (u UserRepository) Get(c context.Context, id primitive.ObjectID) (*user.User, error) {
	var result *user.User
	err := u.worker.GetCollection(collection).FindOne(c, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u UserRepository) GetAll(c context.Context) (*[]user.User, error) {
	cursor, err := u.worker.GetCollection(collection).Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	rst, err := util.DecodeCursor(c, cursor)
	if err != nil {
		return nil, err
	}

	return rst, nil
}

func (u UserRepository) Create(c context.Context, data user.User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	newUser := user.User{
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

	rst, err := u.worker.GetCollection(collection).InsertOne(c, newUser)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if oId, ok := rst.InsertedID.(primitive.ObjectID); ok {
		return oId.Hex(), nil
	} else {
		return "", errors.New("erro ao decodificar o Id")
	}
}

func (u UserRepository) Update(c context.Context, data user.User, id primitive.ObjectID) (*user.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	usr := u.worker.GetCollection(collection).FindOne(c, id)
	if usr == nil {
		fmt.Println("Nenhum usuário encontrado")
		return nil, errors.New("Nenhum usuário encontrado")
	}

	newData := user.User{
		Id:        id,
		Name:      data.Name,
		Email:     data.Email,
		Password:  string(hash),
		State:     data.State,
		Role:      data.Role,
		UserType:  data.UserType,
		CreatedAt: data.CreatedAt,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedBy: "me",
	}

	_, err = u.worker.GetCollection(collection).InsertOne(c, newData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &newData, nil
}

func (u UserRepository) Delete(c context.Context, id primitive.ObjectID) error {
	srst := u.worker.GetCollection(collection).FindOne(c, bson.M{"_id": id})
	if srst == nil {
		return errors.New("Nenhum usuário encontrado")
	}

	_, err := u.worker.GetCollection(collection).DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
