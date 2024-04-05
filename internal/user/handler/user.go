package handler

import (
	"america-rental-backend/internal/user"
	"america-rental-backend/internal/user/ports"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(s ports.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (u UserHandler) Get(g *gin.Context) {
	objId, err := primitive.ObjectIDFromHex(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	usr, err := u.service.Get(g, objId)
	if err != nil {
		g.JSON(http.StatusNoContent, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"user": map[string]interface{}{"data": usr},
	})
}

func (u UserHandler) GetAll(c context.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UserHandler) Create(c context.Context, user user.User) {
	//TODO implement me
	panic("implement me")
}

func (u UserHandler) Update(c context.Context, user user.User, id primitive.ObjectID) {
	//TODO implement me
	panic("implement me")
}

func (u UserHandler) Delete(c context.Context, id primitive.ObjectID) {
	//TODO implement me
	panic("implement me")
}
