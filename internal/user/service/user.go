package service

import (
	"america-rental-backend/internal/user"
	"america-rental-backend/internal/user/ports"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return UserService{repo: repo}
}

func (u UserService) Get(c context.Context, id primitive.ObjectID) (*user.User, error) {
	usr, err := u.repo.Get(c, id)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, errors.New("no user found")
	}

	return usr, nil
}

func (u UserService) GetAll(c context.Context) (*[]user.User, error) {
	users, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserService) Create(c context.Context, user user.User) (*user.User, error) {
	id, err := u.repo.Create(c, user)
	if err != nil {
		return nil, err
	}

	pId, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		return nil, err
	}

	newUser, err := u.Get(c, pId)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u UserService) Update(c context.Context, user user.User, id primitive.ObjectID) (*user.User, error) {
	isPresent, err := u.repo.Get(c, id)
	if err != nil {
		return nil, err
	}

	if isPresent != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hash)

		newUser, err := u.repo.Update(c, user, id)
		if err != nil {
			return nil, err
		}

		return newUser, nil
	} else {
		return nil, errors.New("no user found")
	}
}

func (u UserService) Delete(c context.Context, id primitive.ObjectID) error {
	sRst, err := u.repo.Get(c, id)

	if err != nil {
		return err
	}

	if sRst == nil {
		return errors.New("Nenhum usuário encontrado")
	}

	err = u.repo.Delete(c, id)
	if err != nil {
		return err
	}

	return nil
}
