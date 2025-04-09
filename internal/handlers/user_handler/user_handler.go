package user_handler

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"context"
	"time"

	pb "github.com/banisys/api-gateway/user_service_grpc"
)

func Signup(ginContext *gin.Context) {

	type User struct {
		Name     string
		Email    string `binding:"required`
		Password string `binding:"required`
	}

	var user User

	ginContext.ShouldBindJSON(&user)

	// create grpc client

	//////////////////////////// Mock ////////////////////////////
	flag.Parse()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// invoke signup method from user-service
	response, err := c.Signup(ctx, &pb.UserServiceReq{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		log.Fatalf("user-service error: %v", err)
	}

	fmt.Println("user-service response:", response)

	////////////////////////////////////////////////////////

	ginContext.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}
