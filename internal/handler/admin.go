package handler

import (
	"go-crud/internal/dto/request"
	"go-crud/internal/middleware"
	"go-crud/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	Service service.AdminService
}

func (h *AdminHandler) SignUp(c *gin.Context) {
	var input request.Admin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.SignUp(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AdminHandler) LogIn(c *gin.Context) {
	var input request.Admin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.LogIn(input)
	if err != nil {
		if err.Error() == "authentication failed" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	token, err := middleware.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": token})
}
