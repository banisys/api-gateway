package routes

import (
	"github.com/banisys/api-gateway/internal/grpc_client"
	"github.com/banisys/api-gateway/internal/handlers/user_handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine, GrpcSignupInterface grpc_client.GrpcSignupInterface) {

	route.POST("/signup", func(c *gin.Context) {

		user_handler.Signup(GrpcSignupInterface, c)
	})
}
