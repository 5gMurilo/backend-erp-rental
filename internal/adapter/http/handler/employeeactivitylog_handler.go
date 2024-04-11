package handler

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmployeeActivityLogHandler struct {
	svc ports.EmployeeActivityLogService
}

func NewEmployeeActivityLogHandler(svc ports.EmployeeActivityLogService) EmployeeActivityLogHandler {
	return EmployeeActivityLogHandler{svc}
}

func (h EmployeeActivityLogHandler) Get(g *gin.Context) {
	var empl domain.Employee

	if err := g.ShouldBindJSON(&empl); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	activities, err := h.svc.GetByEmployee(g, empl)
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
