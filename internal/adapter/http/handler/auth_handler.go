package handler

import (
	"america-rental-backend/internal/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc ports.AuthService
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6" minLength:"6"`
}

func NewAuthHandler(svc ports.AuthService) AuthHandler {
	return AuthHandler{svc}
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	authResponse, err := ah.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": authResponse,
	})

}
