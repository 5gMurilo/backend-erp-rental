package handler

import (
	"america-rental-backend/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
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
		_, err := s.svc.SendFile(file, g.PostForm("employee"))
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"erro": err,
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
			"erro": err,
		})
	}

	empname := form.Value["employee"][0]

	files, err := s.svc.ListFiles(empname)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"erro": err,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}
