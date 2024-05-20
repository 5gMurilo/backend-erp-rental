package handler

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeActivityLogHandler struct {
	svc ports.EmployeeActivityLogService
}

func NewEmployeeActivityLogHandler(svc ports.EmployeeActivityLogService) EmployeeActivityLogHandler {
	return EmployeeActivityLogHandler{svc}
}

func (h EmployeeActivityLogHandler) Get(g *gin.Context) {
	id := g.Param("id")

	activities, err := h.svc.GetByEmployee(g, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"data": activities,
	})
}

func (h EmployeeActivityLogHandler) New(g *gin.Context) {
	var newActivity domain.EmployeeActivityLog

	err := g.ShouldBindJSON(&newActivity)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	newActivity.Actor = g.GetString("requestOwner")

	activityLog, err := h.svc.New(g, newActivity)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"activity": activityLog,
	})
}
