package main

import (
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.AuthRoutes(router)
	routes.ProfileRoutes(router)

	router.Run(":8080")
}
