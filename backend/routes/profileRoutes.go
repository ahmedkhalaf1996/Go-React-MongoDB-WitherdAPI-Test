package routes

import (
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/controllers"
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(router *gin.Engine) {
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.GetProfile)
	}
}
