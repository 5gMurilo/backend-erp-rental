package handler

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ReturnableEpiHandler struct {
	svc ports.ReturnableEpiService
}

func NewReturnableEpiHandler(svc ports.ReturnableEpiService) ReturnableEpiHandler {
	return ReturnableEpiHandler{svc}
}

func (r ReturnableEpiHandler) Get(g *gin.Context) {
	id, err := primitive.ObjectIDFromHex(g.Param("id"))
	rst, err := r.svc.GetReturnableEpi(g, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}

	g.JSON(http.StatusOK, rst)
}

func (r ReturnableEpiHandler) GetAll(g *gin.Context) {
	rst, err := r.svc.GetAllReturnableEpi(g)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}

	g.JSON(http.StatusOK, rst)
}

func (r ReturnableEpiHandler) Create(g *gin.Context) {
	var body domain.ReturnableEpi
	err := g.ShouldBindJSON(&body)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}

	rst, err := r.svc.CreateReturnableEpi(g, body)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}

	g.JSON(http.StatusOK, rst)
}

func (r ReturnableEpiHandler) Update(g *gin.Context) {
	id, err := primitive.ObjectIDFromHex(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	var body domain.ReturnableEpi
	err = g.ShouldBindJSON(&body)

	if rst, err := r.svc.UpdateReturnableEpi(g, body, id); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	} else {
		g.JSON(http.StatusOK, rst)
	}
}
func (r ReturnableEpiHandler) Delete(g *gin.Context) {
	id, err := primitive.ObjectIDFromHex(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	err = r.svc.DeleteReturnableEpi(g, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	g.JSON(http.StatusOK, "EPI retornável deletado com sucesso!")
}
