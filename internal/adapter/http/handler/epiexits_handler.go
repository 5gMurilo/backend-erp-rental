package handler

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	g.ShouldBindJSON(&body)

	body.ExitTime = primitive.NewDateTimeFromTime(time.Now())

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
