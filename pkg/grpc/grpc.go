package grpc

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/banisys/api-gateway/user_service_grpc"
)

func GrpcConnetion() (pb.UserServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewUserServiceClient(conn)

	return c, conn
}
