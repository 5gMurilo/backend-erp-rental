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

type EmployeeHandler struct {
	activityService ports.EmployeeActivityLogService
	svc             ports.EmployeeService
}

func NewEmployeeHandler(activityService ports.EmployeeActivityLogService, svc ports.EmployeeService) EmployeeHandler {
	return EmployeeHandler{activityService, svc}
}

func (eh EmployeeHandler) GetAll(g *gin.Context) {
	employees, err := eh.svc.GetAll(g)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, &employees)
}

func (eh EmployeeHandler) GetById(g *gin.Context) {
	id, err := primitive.ObjectIDFromHex(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	emp, err := eh.svc.GetById(g, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
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
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	fmt.Println(g.GetString("requestOwner"))

	rst, err := eh.svc.New(g, newEmployee, g.GetString("requestOwner"))
	if err != nil {
		fmt.Printf("handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	_, err = eh.activityService.New(g, domain.EmployeeActivityLog{
		Activity: "Novo Colaborador",
		Employee: *rst,
		Actor:    g.GetString("requestOwner"),
		At:       primitive.NewDateTimeFromTime(time.Now()),
	})

	if err != nil {
		fmt.Printf("handler error \n%e", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"result": &rst,
	})
}

func (eh EmployeeHandler) Update(g *gin.Context) {
	var data domain.Employee

	oId, err := primitive.ObjectIDFromHex(g.Param("id"))
	if err != nil {
		fmt.Printf("89 handler error \n%s\n", err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	err = g.ShouldBindJSON(&data)
	if err != nil {
		fmt.Printf("98 handler error \n%s\n", err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	newData, err := eh.svc.Update(g, oId, data, g.GetString("requestOwner"))
	if err != nil {
		fmt.Printf("131 handler error \n%s\n", err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	newActivity, err := eh.activityService.New(g, domain.EmployeeActivityLog{
		Activity: "Alteração nos dados do colaborador",
		Employee: *newData,
		Actor:    g.GetString("requestOwner"),
		At:       primitive.NewDateTimeFromTime(time.Now()),
	})

	fmt.Println(newActivity)

	if err != nil {
		fmt.Printf("145 handler error \n%s\n", err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{
			"handler": map[string]interface{}{"error": err.Error()},
		})
		return
	}

	g.JSON(http.StatusOK, &newData)
}
