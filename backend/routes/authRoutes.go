package routes

import (
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
