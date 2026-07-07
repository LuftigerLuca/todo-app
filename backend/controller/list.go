package controller

import (
	"net/http"
	"todo-app/dto"
	"todo-app/service"

	"github.com/gin-gonic/gin"
)

type ListController struct {
	Service service.ListService
}

func (ctrl *ListController) CreateList(c *gin.Context) {
	var req dto.CreateListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	listID, err := ctrl.Service.CreateList(req, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list_id": listID})
}
