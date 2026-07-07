package controller

import (
	"net/http"
	"todo-app/dto"
	"todo-app/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service service.AuthService
}

// @Summary      Register user
// @Description  Creates a new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body  dto.RegisterRequest  true  "Registration data"
// @Success      200  {object}  dto.AuthResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Router       /auth/register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.Service.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}

// @Summary      Login user
// @Description  Authenticates a user and returns a JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body  dto.LoginRequest  true  "Login credentials"
// @Success      200  {object}  dto.AuthResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Router       /auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.Service.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}
