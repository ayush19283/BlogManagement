package router

import (
	"blog-backend/app/api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine) {

	apiGroup := r.Group("/api")
	apiGroup.POST("/signup", controller.SignUp)
	apiGroup.POST("/login", controller.Login)
	apiGroup.POST("/blog", controller.CreateBlog)

	apiGroup.GET("/health", HealthCheck)

}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "API is running",
	})
}
