package controller

import (
	"blog-backend/app/api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BlogRequest struct {
	UserID      int    `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Body        string `json:"body" binding:"required"`
}

func SignUp(c *gin.Context) {
	var req AuthRequest

	log.Printf("Incoming sign-up request: %v", c.Request.Body)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Inavlid Request"})
		return
	}

	user, err := services.SignUp(req.Name, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, user)

}

func Login(c *gin.Context) {
	var req AuthRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Inavlid Request"})
		return
	}

	user, err := services.Login(req.Name, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, user)

}

func CreateBlog(c *gin.Context) {
	var req BlogRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
	}

	blog, err := services.CreateBlog(req.UserID, req.Title, req.Description, req.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, blog)
}
