package handler

import (
	"america-rental-backend/internal/user"
	"america-rental-backend/internal/user/ports"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
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

func (u UserHandler) GetAll(g *gin.Context) {
	rsp, err := u.service.GetAll(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"data": rsp,
	})
}

func (u UserHandler) Create(g *gin.Context) {
	var usr user.User

	if err := g.BindJSON(&usr); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rsp, err := u.service.Create(g, usr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"success": map[string]interface{}{"data": rsp},
	})
}

func (u UserHandler) Update(g *gin.Context) {
	objId, err := primitive.ObjectIDFromHex(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	actualUser, err := u.service.Get(g, objId)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var body user.User
	if err = g.BindJSON(&body); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUserData := user.User{
		Id:            objId,
		Name:          body.Name,
		Email:         body.Email,
		Password:      body.Password,
		State:         body.State,
		Role:          body.Role,
		UserType:      body.UserType,
		MailSignature: body.MailSignature,
		CreatedAt:     actualUser.CreatedAt,
		UpdatedAt:     primitive.NewDateTimeFromTime(time.Now()),
		UpdatedBy:     "Me",
	}

	rst, err := u.service.Update(g, newUserData, objId)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"success": map[string]interface{}{
			"data": rst,
		},
	})
}

func (u UserHandler) Delete(g *gin.Context) {
	objId, err := primitive.ObjectIDFromHex(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = u.service.Delete(g, objId)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"sucesso": "usuário deletado com sucesso!",
	})
}
