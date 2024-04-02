package service

import (
	"america-rental-backend/internal/user"
	"america-rental-backend/internal/user/ports"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo *ports.UserRepository
}

func (u UserService) Get(c *gin.Context, id primitive.ObjectID) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetAll(c *gin.Context) ([]user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) Create(c *gin.Context, user user.User) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) Update(c *gin.Context, user user.User, id primitive.ObjectID) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) Delete(c *gin.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func NewUserService(repo *ports.UserRepository) ports.UserService {
	return UserService{repo: repo}
}
