package middleware

import (
	"america-rental-backend/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token não fornecido",
		})
		return
	}

	_, err := as.service.VerifyToken(strings.ReplaceAll(token, "Bearer ", ""))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token inválido",
		})
		return
	}

	g.Next()
}
