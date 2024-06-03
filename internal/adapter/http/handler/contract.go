package handler

import (
	"america-rental-backend/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContractHandler struct {
	svc ports.ContractService
}

func NewContractHandler(svc ports.ContractService) ContractHandler {
	return ContractHandler{svc}
}

func (c ContractHandler) getContracts(g *gin.Context) {
	contracts, err := c.svc.GetContracts()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"success": contracts,
	})
}
