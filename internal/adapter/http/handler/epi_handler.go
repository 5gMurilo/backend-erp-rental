package handler

import (
	"america-rental-backend/internal/core/ports"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EpiHandler struct {
	svc ports.EpiService
}

func NewEpiHandler(svc ports.EpiService) EpiHandler {
	return EpiHandler{svc}
}

func (e EpiHandler) GetAll(g *gin.Context) {
	rst, err := e.svc.GetAll(g)
	if err != nil {
		fmt.Printf("handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	g.JSON(http.StatusOK, rst)
}

func (e EpiHandler) GetById(g *gin.Context) {
	param := g.Param("id")

	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		fmt.Printf("38 handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}
	rst, err := e.svc.GetEpi(g, id)
	if err != nil {
		fmt.Printf("46 handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	g.JSON(http.StatusOK, rst)
}
