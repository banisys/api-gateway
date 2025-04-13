package main

import (
	"github.com/banisys/api-gateway/internal/grpc_client"
	"github.com/banisys/api-gateway/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	var GrpcSignupStruct grpc_client.GrpcSignupStruct

	routes.RegisterRoutes(route, GrpcSignupStruct)

	route.Run(":8081")
}
