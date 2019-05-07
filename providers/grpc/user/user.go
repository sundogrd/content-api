package user

import (
	"github.com/sundogrd/content-api/grpc_gen/user"
	"google.golang.org/grpc"
	"log"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func NewGrpcUserClient() (user.UserServiceClient, *grpc.ClientConn,error)  {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, nil, err
	}
	c := user.NewUserServiceClient(conn)
	return c, conn, nil
}