package http

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/port"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	srvc port.UserService
}

func NewUserHandler(srvc port.UserService) *UserHandler {
	return &UserHandler{
		srvc: srvc,
	}
}

type registerRequest struct {
	Name          string          `json:"name"`
	Email         string          `json:"email"`
	Password      string          `json:"password"`
	State         string          `json:"state" `
	Role          string          `json:"role" `
	UserType      domain.UserType `json:"userType" `
	MailSignature string          `json:"mail_signature" `
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := domain.User{
		Name:          req.Name,
		Email:         req.Email,
		Password:      req.Password,
		State:         req.State,
		Role:          req.Role,
		UserType:      req.UserType,
		MailSignature: req.MailSignature,
	}

	_, err := handler.srvc.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (handler *UserHandler) GetAllUsers(ctx *gin.Context) {
	res, err := handler.srvc.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

func (handler *UserHandler) GetUser(ctx *gin.Context) {
	ctx.Param("id")
}
