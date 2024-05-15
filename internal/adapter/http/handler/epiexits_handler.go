package handler

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EpiExitsHandler struct {
	svc ports.EpiExitsService
}

func NewEpiExitsHandler(svc ports.EpiExitsService) EpiExitsHandler {
	return EpiExitsHandler{svc}
}

func (e EpiExitsHandler) GetAll(g *gin.Context) {
	rst, err := e.svc.GetExits(g)
	if err != nil {
		fmt.Printf("21 - handler - error: %s\n", err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, rst)
}

func (e EpiExitsHandler) New(g *gin.Context) {
	var body domain.EpiExits

	err := g.ShouldBindJSON(&body)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	body.GaveBy = g.GetString("requestOwner")

	rst, err := e.svc.NewExit(g, body)
	if err != nil {
		fmt.Printf("34 - handler - error: %s\n", err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, &rst)
}
