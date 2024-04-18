package handler

import (
	"america-rental-backend/internal/core/domain"
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

func (e EpiHandler) NewEpi(g *gin.Context) {
	userType := g.GetString("userType")

	if userType != "ADMIN" {
		g.JSON(http.StatusUnauthorized, gin.H{
			"handler": map[string]interface{}{"error": "Usuário não autorizado a realizar esse tipo de requisição"},
		})
		return
	}

	var body domain.Epi

	err := g.ShouldBindJSON(&body)
	if err != nil {
		fmt.Printf("71 handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	rst, err := e.svc.NewEpi(g, body)
	if err != nil {
		fmt.Printf("80 handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	g.JSON(http.StatusCreated, rst)
}

func (e EpiHandler) DeleteEpi(g *gin.Context) {
	userType := g.GetString("userType")

	if userType != "ADMIN" {
		g.JSON(http.StatusUnauthorized, gin.H{
			"handler": map[string]interface{}{"error": "Usuário não autorizado a realizar esse tipo de requisição"},
		})
		return
	}

	param := g.Param("id")

	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		fmt.Printf("104 handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	err = e.svc.DeleteEpi(g, id)
	if err != nil {
		fmt.Printf("113 handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	g.JSON(http.StatusOK, "EPI Removido da base de dados")
}

func (e EpiHandler) UpdateEpi(g *gin.Context) {
	userType := g.GetString("userType")

	if userType != "ADMIN" {
		g.JSON(http.StatusUnauthorized, gin.H{
			"handler": map[string]interface{}{"error": "Usuário não autorizado a realizar esse tipo de requisição"},
		})
		return
	}
	var body domain.Epi
	err := g.ShouldBindJSON(&body)
	if err != nil {
		fmt.Printf("135 handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	param := g.Param("id")

	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		fmt.Printf("146 handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	rst, err := e.svc.UpdateEpi(g, id, body)
	if err != nil {
		fmt.Printf("156 handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	g.JSON(http.StatusOK, rst)
}
