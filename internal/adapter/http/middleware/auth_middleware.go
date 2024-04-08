package middleware

import (
	"america-rental-backend/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct {
	service ports.TokenService
}

func NewAuthMiddleware(as ports.TokenService) *AuthMiddleware {
	return &AuthMiddleware{
		service: as,
	}
}

func (as AuthMiddleware) AuthenticationMiddleware(g *gin.Context) {
	token := g.GetHeader("Authorization")
	if token == "" {
		g.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token não fornecido",
		})
		return
	}

	_, err := as.service.VerifyToken(token)
	if err != nil {
		return
	}

	g.Next()
}
