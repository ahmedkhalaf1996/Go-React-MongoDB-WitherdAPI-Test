package main

import (
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Type", "Accept"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	routes.AuthRoutes(router)
	routes.ProfileRoutes(router)

	router.Run(":8080")
}
