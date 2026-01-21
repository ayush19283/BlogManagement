package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Name    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context){
	var req AuthRequest

	err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Inavlid Request"})
		return
	}

}