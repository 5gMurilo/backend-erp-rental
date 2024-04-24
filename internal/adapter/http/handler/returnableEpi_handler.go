package handler

import "america-rental-backend/internal/core/ports"

type ReturnableEpiHandler struct {
	svc ports.ReturnableEpiService
}

func NewReturnableEpiService(svc ports.ReturnableEpiService) ReturnableEpiHandler {
	return ReturnableEpiHandler{svc}
}
