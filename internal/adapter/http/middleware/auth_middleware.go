package middleware

import (
	"america-rental-backend/internal/core/ports"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

	payload, err := as.service.VerifyToken(strings.ReplaceAll(token, "Bearer ", ""))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token inválido",
		})
		return
	}

	g.Set("requestOwner", payload.Name)

	g.Next()
}
