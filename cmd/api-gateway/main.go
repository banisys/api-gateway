package main

import (
	"github.com/banisys/api-gateway/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	routes.RegisterUserRoutes(route)

	route.Run()
}
