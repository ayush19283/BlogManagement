package controller

import (
	"blog-backend/app/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Name     string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context) {
	var req AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Inavlid Request"})
		return
	}

	user, err := services.SignUp(req.Name, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, user)

}
