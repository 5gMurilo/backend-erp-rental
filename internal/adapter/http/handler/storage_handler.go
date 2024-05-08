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
		_, err := s.svc.SendFile(file, g.PostForm("employee"), g.PostForm("type"), g.GetString("requestOwner"))
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
	employee := g.Param("name")
	files, err := s.svc.ListFiles(employee)
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

func (s StorageHandler) Delete(g *gin.Context) {
	driveItemId := g.Param("driveItem")

	err := s.svc.DeleteFile(driveItemId)
	if err != nil {
		g.JSON(500, gin.H{
			"erro": err.Error(),
		})
		return
	}

	g.JSON(200, gin.H{
		"sucesso": "O arquivo foi removido do onedrive com sucesso!",
	})
}
