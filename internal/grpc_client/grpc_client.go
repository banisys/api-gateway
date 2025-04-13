package grpc_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/banisys/api-gateway/internal/models"
	"github.com/banisys/api-gateway/pkg/grpc"

	pb "github.com/banisys/api-gateway/user_service_grpc"
)

type GrpcSignupInterface interface {
	GrpcSignup(models.User) *pb.UserServiceRes
}

type GrpcSignupStruct struct {
}

func (g GrpcSignupStruct) GrpcSignup(user models.User) *pb.UserServiceRes {

	c, conn := grpc.GrpcConnetion()

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// invoke signup method from user-service
	response, err := c.Signup(ctx, &pb.UserServiceReq{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	fmt.Println(response)

	if err != nil {
		log.Fatalf("user-service error: %v", err)
	}

	fmt.Println("user-service response:", response.Message)

	return response
}
