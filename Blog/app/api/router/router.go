package router

import (
	"blog-backend/app/api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine) {

	apiGroup := r.Group("/api")

	userGroup := apiGroup.Group("/users")
	{
		userGroup.POST("/singup", controller.SignUp)

	}

}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "API is running",
	})
}
