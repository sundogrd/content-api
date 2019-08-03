package content

import (
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-api/grpc_gen/content"
	grpcUtils "github.com/sundogrd/gopkg/grpc"
	"google.golang.org/grpc"
)

const (
	//address     = "localhost:50052"
	address     = "sundog.content"
	defaultName = "sundog.content"
)

// NewGrpcContentClient ...
func NewGrpcContentClient() (content.ContentServiceClient, *grpc.ClientConn, error) {
	r, err := grpcUtils.NewGrpcResolover()
	if err != nil {
		return nil, nil, err
	}
	b := grpc.RoundRobin(r)
	// Set up a connection to the server.
	// WithBlock https://gocn.vip/question/931  wtf
	conn, err := grpc.Dial(address, grpc.WithBalancer(b), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
		return nil, nil, err
	}
	c := content.NewContentServiceClient(conn)
	return c, conn, nil
}
