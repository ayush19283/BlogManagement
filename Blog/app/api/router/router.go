package router

import (
	"github.com/gin-gonic/gin"
)

func setupRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.POST("/signup")

}
