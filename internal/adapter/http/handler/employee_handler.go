package handler

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeHandler struct {
	svc ports.EmployeeService
}

func NewEmployeeHandler(svc ports.EmployeeService) EmployeeHandler {
	return EmployeeHandler{svc}
}

func (eh EmployeeHandler) GetAll(g *gin.Context) {
	employees, err := eh.svc.GetAll(g)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"result": &employees,
	})
}

func (eh EmployeeHandler) GetById(g *gin.Context) {
	id, err := primitive.ObjectIDFromHex(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	emp, err := eh.svc.GetById(g, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"result": &emp,
	})
}

func (eh EmployeeHandler) New(g *gin.Context) {
	var newEmployee domain.Employee

	err := g.ShouldBindJSON(&newEmployee)
	if err != nil {
		fmt.Printf("handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err},
		})
		return
	}

	fmt.Println(g.GetString("requestOwner"))

	rst, err := eh.svc.New(g, newEmployee, g.GetString("requestOwner"))
	if err != nil {
		fmt.Printf("handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err},
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"result": &rst,
	})
}

func (eh EmployeeHandler) Update(g *gin.Context) {
	var data domain.Employee

	fmt.Print("h")

	err := g.ShouldBindJSON(&data)
	if err != nil {
		fmt.Printf("handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err},
		})
		return
	}

	oId, err := primitive.ObjectIDFromHex(g.Param("id"))
	if err != nil {
		fmt.Printf("handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err},
		})
		return
	}

	newData, err := eh.svc.Update(g, oId, data, g.GetString("requestOwner"))
	if err != nil {
		fmt.Printf("handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err},
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"result": &newData,
	})
}
