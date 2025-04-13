package user_handler

import (
	"net/http"

	"github.com/banisys/api-gateway/internal/grpc_client"
	"github.com/banisys/api-gateway/internal/models"
	"github.com/gin-gonic/gin"
)

func Signup(GrpcSignupInterface grpc_client.GrpcSignupInterface, c *gin.Context) {

	var user models.User
	c.ShouldBindJSON(&user)

	grpcRes := GrpcSignupInterface.GrpcSignup(user)

	c.JSON(http.StatusCreated, gin.H{
		"message": grpcRes.Message,
		"user_id": grpcRes.UserId,
	})

}
