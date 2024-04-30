package handler

import (
	"america-rental-backend/internal/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StorageHandler struct {
	svc ports.StorageServices
}

func NewStorageHandler(svc ports.StorageServices) StorageHandler {
	return StorageHandler{svc}
}

func (s StorageHandler) Create(g *gin.Context) {

	form, err := g.MultipartForm()
	if err != nil {
		return
	}

	files := form.File["files"]
	for _, file := range files {
		_, err := s.svc.SendFile(file, g.PostForm("employee"), g.PostForm("employee"), g.GetString("requestOwner"))
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})
			return
		}
	}

	g.JSON(http.StatusOK, gin.H{
		"sucesso": "arquivos enviados com sucesso!",
	})
}

func (s StorageHandler) List(g *gin.Context) {
	form, err := g.MultipartForm()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
	}

	empName := form.Value["employee"][0]

	files, err := s.svc.ListFiles(empName)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}
